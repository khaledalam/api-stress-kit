# 🚀 api-stress-kit

**api-stress-kit** is a minimal, production-ready starter to **build**, **benchmark**, and **stress-test** REST APIs using **Go**, **Gin**, and popular load-testing tools — **Vegeta** and **k6**.

---

## ✨ Features
- ⚡ Fast HTTP API built with **Gin**
- 🧪 Stress testing using **Vegeta** & **k6**
- 🐳 Dockerfile + Compose for reproducible environments
- ✅ Basic unit test for `/ping`
- 🧩 Ready for profiling with `pprof`

---

## ⚙️ Quick Start

### Run Locally
```bash
go mod tidy
go run main.go
```

### Install Load Testing Tools
```bash
# Vegeta
go install github.com/tsenart/vegeta/v12@latest
# or using Homebrew
brew install vegeta

# k6
brew install k6
```

---

## 🔥 Stress Testing Examples

### Vegeta
```bash
echo "GET http://localhost:8080/ping" | vegeta attack -duration=15s -rate=200 | vegeta report
```

### k6
```bash
k6 run tests/stress/k6_test.js
```

---

## 🐳 Run with Docker
```bash
docker-compose up --build
# Server available at: http://localhost:8080
```

---

## 🧠 Endpoints
| Method | Path | Description | Example Response |
|--------|------|--------------|------------------|
| `GET` | `/ping` | Health check with latency info | `{ "message": "pong", "latency_ms": <int> }` |
| `GET` | `/health` | Simple status check | `{ "status": "ok" }` |
| `GET` | `/users` | List of demo users | `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]` |

---

## 📊 Output

<details>
<summary><b>▶️ Click to expand full test results (k6 + Vegeta)</b></summary>

### 🧪 k6 Test Output
```bash
% k6 run tests/stress/k6_test.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: tests/stress/k6_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 20 max VUs, 50s max duration (incl. graceful stop):
              * default: 20 looping VUs for 20s (gracefulStop: 30s)


  █ TOTAL RESULTS 

    checks_total.......: 800     39.43/s
    checks_succeeded...: 100.00% 800 out of 800
    checks_failed......: 0.00%   0 out of 800

    ✓ status was 200

    HTTP
    http_req_duration..............: avg=5.71ms   min=710µs    med=5.67ms   max=15.08ms  
      { expected_response:true }...: avg=5.71ms   min=710µs    med=5.67ms   max=15.08ms  
    http_req_failed................: 0.00%  0 out of 800
    http_reqs......................: 800    39.43/s

    EXECUTION
    iteration_duration.............: avg=507.21ms min=500.91ms med=507.06ms max=519.43ms
    iterations.....................: 800    39.43/s
    vus............................: 20     min=20       max=20
    vus_max........................: 20     min=20       max=20

    NETWORK
    data_received..................: 125 kB 6.2 kB/s
    data_sent......................: 59 kB  2.9 kB/s

running (20.3s), 00/20 VUs, 800 complete and 0 interrupted iterations
default ✓ [======================================] 20 VUs  20s
```

---

### ⚡ Vegeta Test Output
```bash
% echo "GET http://localhost:8080/users" | vegeta attack -duration=15s -rate=200 | vegeta report

Requests      [total, rate, throughput]         3000, 200.07, 200.05
Duration      [total, attack, wait]             14.996s, 14.995s, 1.339ms
Latencies     [min, mean, 50, 90, 95, 99, max]  428.25µs, 1.749ms, 1.252ms, 3.208ms, 4.227ms, 8.217ms, 33.845ms
Bytes In      [total, mean]                     141000, 47.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:3000  
Error Set:
```
</details>

---

## ⚙️ Performance Comparison

Command used:
```bash
echo "GET http://localhost:8080/users" | vegeta attack -duration=10s -rate=2000 | vegeta report
```

---

### 🖥️ Native Execution (Direct Go Run)
```bash
go run main.go
```

#### 🕒 Idle (0 Requests)
<p align="center">
  <img src="docs/os-0.png" width="70%" alt="Native - Idle CPU Graph" />
</p>

#### 🚀 Load Test (10s @ 2000 RPS)
<p align="center">
  <img src="docs/os-cpu.png" width="70%" alt="Native - CPU Load" />
  <br/>
  <img src="docs/os-stress.png" width="70%" alt="Native - Stress Graph" />
</p>

---

### 🐳 Inside Docker (Containerized Execution)
```bash
docker-compose up --build
```

#### 🕒 Idle (0 Requests)
<p align="center">
  <img src="docs/docker-0.png" width="70%" alt="Docker - Idle CPU Graph" />
</p>

#### 🚀 Load Test (10s @ 2000 RPS)
<p align="center">
  <img src="docs/docker-cpu.png" width="70%" alt="Docker - CPU Load" />
  <br/>
  <img src="docs/docker-stress.png" width="70%" alt="Docker - Stress Graph" />
</p>

---

### 📊 Summary
| Environment | Command | Overhead Observed | Network Mode |
|--------------|----------|-------------------|---------------|
| 🖥️ Native | `go run main.go` | Minimal CPU overhead | Direct host |
| 🐳 Docker | `docker-compose up --build` | Slightly higher baseline CPU (due to VM bridge) | Docker bridge |

> 🧠 *Observation:* Running under Docker introduces an extra CPU overhead layer caused by virtualization, filesystem overlay, and bridged networking — especially visible during high RPS bursts.

---

## 🧾 License
**MIT License**

---

## 👤 Maintainer
**Khaled Alam**  
Full-Stack Software Engineer 

🌐 [Portfolio](https://khaledalam.net)  
✉️ [khaledalam.net@gmail.com](mailto:khaledalam.net@gmail.com)  
💼 [LinkedIn](https://linkedin.com/in/khaledalam)

