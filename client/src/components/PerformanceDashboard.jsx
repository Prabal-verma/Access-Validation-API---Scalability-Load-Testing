// src/components/PerformanceDashboard.jsx
import React, { useState, useEffect } from 'react';
import { Line } from 'react-chartjs-2';

const PerformanceDashboard = () => {
  const [metrics, setMetrics] = useState({
    requestsPerSecond: 0,
    averageLatency: 0,
    errorRate: 0
  });

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl mb-4">Performance Metrics</h1>
      
      <div className="grid grid-cols-3 gap-4 mb-8">
        <div className="bg-white p-4 rounded shadow">
          <h3 className="text-gray-500">Requests/Second</h3>
          <p className="text-2xl">{metrics.requestsPerSecond}</p>
        </div>
        <div className="bg-white p-4 rounded shadow">
          <h3 className="text-gray-500">Avg Latency (ms)</h3>
          <p className="text-2xl">{metrics.averageLatency}</p>
        </div>
        <div className="bg-white p-4 rounded shadow">
          <h3 className="text-gray-500">Error Rate</h3>
          <p className="text-2xl">{metrics.errorRate}%</p>
        </div>
      </div>

      {/* Add charts for historical data */}
    </div>
  );
};

export default PerformanceDashboard;