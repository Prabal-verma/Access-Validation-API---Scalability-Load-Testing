// src/components/TestConsole.jsx
import React, { useState } from 'react';

const TestConsole = () => {
  const [request, setRequest] = useState({
    game_id: '',
    country: '',
    app_version: '',
    platform: '',
    app_type: ''
  });
  const [response, setResponse] = useState(null);

  const testAccess = async () => {
    try {
      const res = await fetch('http://localhost:8080/validate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(request)
      });
      const data = await res.json();
      setResponse(data);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl mb-4">Access Validation Test Console</h1>
      
      <div className="grid grid-cols-2 gap-8">
        {/* Request Panel */}
        <div className="bg-white p-4 rounded shadow">
          <h2 className="text-xl mb-2">Request</h2>
          <div className="space-y-4">
            <input
              placeholder="Game ID"
              value={request.game_id}
              onChange={e => setRequest({...request, game_id: e.target.value})}
              className="border p-2 rounded w-full"
            />
            <input
              placeholder="Country"
              value={request.country}
              onChange={e => setRequest({...request, country: e.target.value})}
              className="border p-2 rounded w-full"
            />
            <input
              placeholder="App Version"
              value={request.app_version}
              onChange={e => setRequest({...request, app_version: e.target.value})}
              className="border p-2 rounded w-full"
            />
            <select
              value={request.platform}
              onChange={e => setRequest({...request, platform: e.target.value})}
              className="border p-2 rounded w-full"
            >
              <option value="">Select Platform</option>
              <option value="iOS">iOS</option>
              <option value="Android">Android</option>
            </select>
            <select
              value={request.app_type}
              onChange={e => setRequest({...request, app_type: e.target.value})}
              className="border p-2 rounded w-full"
            >
              <option value="">Select App Type</option>
              <option value="mobile">Mobile</option>
              <option value="desktop">Desktop</option>
            </select>
            <button
              onClick={testAccess}
              className="bg-blue-500 text-white px-4 py-2 rounded w-full"
            >
              Test Access
            </button>
          </div>
        </div>

        {/* Response Panel */}
        <div className="bg-white p-4 rounded shadow">
          <h2 className="text-xl mb-2">Response</h2>
          {response && (
            <div className={`p-4 rounded ${response.allowed ? 'bg-green-100' : 'bg-red-100'}`}>
              <p className="text-lg font-semibold">
                Access: {response.allowed ? 'Allowed' : 'Denied'}
              </p>
              {response.reason && (
                <p className="text-gray-600 mt-2">Reason: {response.reason}</p>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default TestConsole;