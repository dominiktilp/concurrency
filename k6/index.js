import {group, sleep} from 'k6';
import http from 'k6/http';
import { Rate, Trend } from 'k6/metrics';


const successRate = new Rate("Requests Success rate");
const duration = new Trend("Requests Duration");

const URL = __ENV.URL;

export let options = {
  stages: [
    { duration: "10s", target: 100 },
    { duration: "30s", target: 0 }    
  ],
  noConnectionReuse: true,
  rps: 500,
  batch: 10,
  userAgent: "MyK6UserAgentString/1.0"
};

function testUrl(url) {
  const response = http.get(url);
  duration.add(response.timings.duration);
  successRate.add(response.status === 200 ? 1 : 0);
}

export default function() {
  testUrl(URL);
};
