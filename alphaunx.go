package main

import (
    "bufio"
    "crypto/tls"
    "fmt"
    "net"
    "net/http"
    "os"
    "strings"
    "sync"
    "time"
    "sync/atomic"
)

const (
    Reset     = "\033[0m"
    Red       = "\033[31m"
    Green     = "\033[32m"
    Yellow    = "\033[33m"
    Gray      = "\033[37m"
)

const banner = Green + `
▄▀█ █░░ █▀█ █░█ ▄▀█ █▀ █▀▀ ▄▀█ █▄░█
█▀█ █▄▄ █▀▀ █▀█ █▀█ ▄█ █▄▄ █▀█ █░▀█
              v1.0
[*]Developed by Dark503

[*] Beta: Alphaunx-Go Scanner
` + Reset

const usage = Green + `
[*] Telegram Channel: @alphaunx
` + Reset

type ServerInfo struct {
    Type    string
    Version string
}

type Result struct {
    Domain     string
    StatusCode int
    Server     ServerInfo
    Error      error
    Duration   time.Duration
}

type Scanner struct {
    client           *http.Client
    results          []Result
    mu              sync.Mutex
    startTime       time.Time
    totalDomains    int
    processedDomains int64
}

func NewScanner(totalDomains int) *Scanner {
    transport := &http.Transport{
        DialContext: (&net.Dialer{
            Timeout:   5 * time.Second,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
        MaxIdleConns:       100,
        IdleConnTimeout:    90 * time.Second,
        DisableCompression: true,
    }

    return &Scanner{
        client: &http.Client{
            Transport: transport,
            Timeout:   10 * time.Second,
        },
        startTime:    time.Now(),
        totalDomains: totalDomains,
        results:      make([]Result, 0, totalDomains),
    }
}

func (s *Scanner) processDomain(domain string) Result {
    if !strings.HasPrefix(domain, "http") {
        domain = "http://" + domain
    }

    start := time.Now()
    
    // Try HTTP first
    resp, err := s.client.Head(domain)
    if err != nil {
        // If HTTP fails, try HTTPS
        if !strings.HasPrefix(domain, "https") {
            httpsURL := "https://" + strings.TrimPrefix(domain, "http://")
            resp2, err2 := s.client.Head(httpsURL)
            if err2 == nil {
                resp = resp2
                err = nil
                domain = httpsURL
            }
        }
    }
    
    duration := time.Since(start)

    result := Result{
        Domain:   domain,
        Duration: duration,
    }

    if err != nil {
        result.Error = err
        return result
    }
    defer resp.Body.Close()

    result.StatusCode = resp.StatusCode
    result.Server = detectServer(resp)
    return result
}

func (s *Scanner) addResult(result Result) {
    s.mu.Lock()
    s.results = append(s.results, result)
    s.mu.Unlock()
    atomic.AddInt64(&s.processedDomains, 1)
    
    percentage := float64(atomic.LoadInt64(&s.processedDomains)) / float64(s.totalDomains) * 100
    
    // Consider a domain working if it responds with any status code
    isWorking := result.Error == nil
    
    if !isWorking {
        fmt.Printf("\r%s[%6.1f%%] %-40s %s✗%s%s\n",
            Gray, percentage, result.Domain, Red, Reset, Gray)
        return
    }

    fmt.Printf("\r%s[%6.1f%%] %-40s %s✓%s %s[%d]%s %s%s%s\n",
        Gray, percentage,
        result.Domain,
        Green, Reset,
        Yellow, result.StatusCode, Gray,
        Yellow, result.Server.Type, Reset)
}

func detectServer(resp *http.Response) ServerInfo {
    server := ServerInfo{Type: "Unknown", Version: ""}
    
    serverHeader := resp.Header.Get("Server")
    if serverHeader != "" {
        serverLower := strings.ToLower(serverHeader)
        switch {
        case strings.Contains(serverLower, "nginx"):
            server.Type = "Nginx"
        case strings.Contains(serverLower, "apache"):
            server.Type = "Apache"
        case strings.Contains(serverLower, "microsoft-iis"):
            server.Type = "IIS"
        case strings.Contains(serverLower, "litespeed"):
            server.Type = "LiteSpeed"
        case strings.Contains(serverLower, "caddy"):
            server.Type = "Caddy"
        case strings.Contains(serverLower, "tomcat"):
            server.Type = "Tomcat"
        case strings.Contains(serverLower, "nodejs"):
            server.Type = "Node.js"
        case strings.Contains(serverLower, "jetty"):
            server.Type = "Jetty"
        case strings.Contains(serverLower, "gunicorn"):
            server.Type = "Gunicorn"
        case strings.Contains(serverLower, "express"):
            server.Type = "Express"
        case strings.Contains(serverLower, "lighttpd"):
            server.Type = "Lighttpd"
        case strings.Contains(serverLower, "cherokee"):
            server.Type = "Cherokee"
        case strings.Contains(serverLower, "openresty"):
            server.Type = "OpenResty"
        case strings.Contains(serverLower, "google frontend"):
            server.Type = "Google Frontend"
        case strings.Contains(serverLower, "cloudflare"):
            server.Type = "Cloudflare"
        case strings.Contains(serverLower, "cloudfront"):
            server.Type = "CloudFront"
        case strings.Contains(serverLower, "fastly"):
            server.Type = "Fastly"
        case strings.Contains(serverLower, "akamai"):
            server.Type = "Akamai"
        case strings.Contains(serverLower, "aws"):
            server.Type = "AWS"
        case strings.Contains(serverLower, "azure"):
            server.Type = "Azure"
        case strings.Contains(serverLower, "heroku"):
            server.Type = "Heroku"
        case strings.Contains(serverLower, "vercel"):
            server.Type = "Vercel"
        case strings.Contains(serverLower, "netlify"):
            server.Type = "Netlify"
        case strings.Contains(serverLower, "kestrel"):
            server.Type = "Kestrel"
        case strings.Contains(serverLower, "traefik"):
            server.Type = "Traefik"
        case strings.Contains(serverLower, "envoy"):
            server.Type = "Envoy"
        case strings.Contains(serverLower, "istio"):
            server.Type = "Istio"
        case strings.Contains(serverLower, "kong"):
            server.Type = "Kong"
        case strings.Contains(serverLower, "haproxy"):
            server.Type = "HAProxy"
        case strings.Contains(serverLower, "varnish"):
            server.Type = "Varnish"
        case strings.Contains(serverLower, "squid"):
            server.Type = "Squid"
        case strings.Contains(serverLower, "iplanet"):
            server.Type = "iPlanet"
        case strings.Contains(serverLower, "oracle"):
            server.Type = "Oracle"
        case strings.Contains(serverLower, "ibm"):
            server.Type = "IBM"
        case strings.Contains(serverLower, "sun"):
            server.Type = "Sun"
        case strings.Contains(serverLower, "zeus"):
            server.Type = "Zeus"
        case strings.Contains(serverLower, "booster"):
            server.Type = "Booster"
        case strings.Contains(serverLower, "ats"):
            server.Type = "ATS"
        case strings.Contains(serverLower, "tengine"):
            server.Type = "Tengine"
        case strings.Contains(serverLower, "resin"):
            server.Type = "Resin"
        case strings.Contains(serverLower, "glassfish"):
            server.Type = "GlassFish"
        case strings.Contains(serverLower, "weblogic"):
            server.Type = "WebLogic"
        case strings.Contains(serverLower, "websphere"):
            server.Type = "WebSphere"
        default:
            server.Type = serverHeader
        }
    }

    poweredBy := resp.Header.Get("X-Powered-By")
    if poweredBy != "" {
        if server.Type == "Unknown" {
            server.Type = poweredBy
        } else {
            server.Version = poweredBy
        }
    }

    return server
}

func main() {
    if len(os.Args) != 2 {
        fmt.Print(banner)
        fmt.Print(usage)
        os.Exit(1)
    }

    hostfile := os.Args[1]
    file, err := os.Open(hostfile)
    if err != nil {
        fmt.Printf("%s[!] Error: Cannot open file %s: %v%s\n", Green, hostfile, err, Reset)
        os.Exit(1)
    }
    defer file.Close()

    fmt.Print(banner)
    fmt.Printf("%s[+] Reading domains from: %s%s\n", Green, hostfile, Reset)

    var domains []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            domains = append(domains, line)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("%s[!] Error reading file: %v%s\n", Green, err, Reset)
        os.Exit(1)
    }

    totalDomains := len(domains)
    if totalDomains == 0 {
        fmt.Printf("%s[!] No domains found in file%s\n", Green, Reset)
        os.Exit(1)
    }

    fmt.Printf("%s[+] Found %d domains%s\n", Green, totalDomains, Reset)

    var workers int
    for {
        fmt.Printf("%s[+] Enter number of workers (10-500): %s", Green, Reset)
        _, err := fmt.Scan(&workers)
        if err != nil || workers < 10 || workers > 500 {
            fmt.Printf("%s[!] Invalid input. Please enter a number between 10 and 500.%s\n", Green, Reset)
        } else {
            break
        }
    }

    fmt.Printf("%s[+] Starting scan with %d workers%s\n\n", Green, workers, Reset)

    s := NewScanner(totalDomains)
    semaphore := make(chan struct{}, workers)
    var wg sync.WaitGroup

    for _, domain := range domains {
        wg.Add(1)
        go func(d string) {
            defer wg.Done()
            semaphore <- struct{}{}
            result := s.processDomain(d)
            s.addResult(result)
            <-semaphore
        }(domain)
    }

    wg.Wait()

    fmt.Printf("\n%s[+] Scan completed in %.2f seconds%s\n", 
        Green, time.Since(s.startTime).Seconds(), Reset)
    fmt.Printf("%s[+] Results (only successful responses):%s\n", Green, Reset)
    for _, result := range s.results {
        if result.Error == nil { // Show all responding domains
            cleanDomain := strings.TrimPrefix(result.Domain, "http://")
            cleanDomain = strings.TrimPrefix(cleanDomain, "https://")
            fmt.Printf("%s%s%s\n", 
                Green, cleanDomain, Reset)
        }
    }

    fmt.Printf("\n%s[+] Telegram Channel: @alphaunx  Owner: @l300e %s\n", Green, Reset)
}
