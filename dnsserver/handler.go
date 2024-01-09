package dnsserver

import (
	"errors"
	"log"
	"net"

	"github.com/miekg/dns"
	"github.com/op-y/gdns/cache"
	"github.com/op-y/gdns/storage"
)

var (
	ProxyQueryFailedErr = errors.New("all proxy query failed")
)

type DefaultHandler struct {
	nameservers []string
	timeout     int
	ttl         int
}

func (dh *DefaultHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	dc := cache.Default()
	ds := storage.Default()

	name := r.Question[0].Name
	domain := name
	if domain[len(domain)-1] == '.' {
		domain = domain[0 : len(domain)-1]
	}

	// records from cache
	records, err := dc.QueryA(domain)
	if err == nil && len(records) > 0 {
		m := new(dns.Msg)
		m.SetReply(r)
		m.Authoritative = true
		dnsRR := []dns.RR{}
		for _, record := range records {
			rr := new(dns.A)
			rr.Hdr = dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(dh.ttl)}
			rr.A = net.ParseIP(record)
			dnsRR = append(dnsRR, rr)
		}
		m.Answer = dnsRR
		w.WriteMsg(m)
		return
	}

	// records from storage
	records, err = ds.QueryA(domain)
	if err == nil && len(records) > 0 {
		if err := dc.UpdateA(domain, records); err != nil {
			log.Printf("update cache failed %s\n", err.Error())
		}

		m := new(dns.Msg)
		m.SetReply(r)
		m.Authoritative = true
		var dnsRR []dns.RR

		for _, record := range records {
			rr := new(dns.A)
			rr.Hdr = dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(dh.ttl)}
			rr.A = net.ParseIP(record)
			dnsRR = append(dnsRR, rr)
		}
		m.Answer = dnsRR
		w.WriteMsg(m)
		return
	}

	// search failed
	if err != nil {
		log.Printf("seach from storage failed: %s", err.Error())
		dh.HandleFailed(w, r)
		return
	}

	proxyMsg, err := dh.Proxy(r)
	if err != nil {
		log.Printf("upstream resovle failed: %s", err.Error())
		dh.HandleFailed(w, r)
		return
	}
	// just cache proxy response
	//newRecords := []string{}
	//for _, ans := range proxyMsg.Answer {
	//	if ans.Header().Rrtype == dns.TypeA {
	//		rfields := strings.Split(ans.String(), "\t")
	//		address := rfields[len(rfields)-1]
	//		newRecords = append(newRecords, address)
	//	}
	//}
	//if err := dc.UpdateA(domain, newRecords); err != nil {
	//	log.Printf("update cache failed %s\n", err.Error())
	//}
	w.WriteMsg(proxyMsg)
	return
}

func (dh *DefaultHandler) Proxy(msg *dns.Msg) (*dns.Msg, error) {
	for _, ns := range dh.nameservers {
		r, err := dns.Exchange(msg, ns)
		if err != nil {
			continue
		}
		if r.Answer == nil {
			continue
		}
		return r, nil
	}
	return nil, ProxyQueryFailedErr
}

func (dh *DefaultHandler) HandleFailed(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetRcode(r, dns.RcodeServerFailure)
	w.WriteMsg(m)
	return
}
