# Monitoring and Metrics Collection Automation
End to end build of go app, docker container, prometheus, and grafana with dashboard.

## How to run the stack

1. install docker - https://docs.docker.com/get-docker/

   install docker compose - https://docs.docker.com/compose/install/
   
   install git https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
   
2. clone repsitory
   e.g git clone https://github.com/r00t4ccess/monitoring_automation.git

3. cd monitoring_automation 

4. docker-compose up

   Once everything is loaded the screen will output log information

   To kill press ctrl+c

## Accessing the components

- Api endpoint - http://localhost
- App metrics endpoint - http://localhost:9100
- Prometheus web interface - http://localhost:9090
- Grafana web interface - http://localhost:3000

#### Grafana credentials
- user: admin 
- password: password1234

## Api Server
The api can be accessed from a browser or the command like with something like curl

```sh
user $: curl localhost

[{"Phrase":"Hello World"}]
```
The /metrics endpoint can be accessed the same way

```
user $: curl localhost:9100/metrics

# HELP hello_world_http_counter The total number of requests to the api
# TYPE hello_world_http_counter counter
hello_world_http_counter 1
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.47
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 11
```

## Prometheus scraper
The configuration is found in
prometheus/prometheus.yml

```
  # Scrape the Node Exporter every 5 seconds.
  - job_name: 'node'
    scrape_interval: 5s
    static_configs:
      - targets: ['apiserver:9100']
```

## Grafana Chart

The Grafana chart pulls the value "hello_world_http_counter" from the prometheus data set
and displays the results in a graph

![grafana screen grab](https://github.com/r00t4ccess/monitoring_automation/blob/master/images/helloapi_grafana.png?raw=true)

### Container network
The containers talk to each other over the internal docker network called "internal"

The internal dns records are
- apiserver
- prometheus
- grafana
