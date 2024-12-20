import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import RulesDashboard from './components/RulesDashboard';
import TestConsole from './components/TestConsole';
import PerformanceDashboard from './components/PerformanceDashboard';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-blue-200">
        <nav className="bg-gray-300 shadow-lg">
          <div className="max-w-7xl mx-auto px-4">
            <div className="flex justify-between h-16">
              <div className="flex">
                <div className="flex-shrink-0 flex items-center">
                  <h1 className="text-xl font-bold">Game Auth Service</h1>
                </div>
                <div className="ml-6 flex space-x-8">
                  <Link to="/admin" className="text-gray-500 hover:text-gray-700 px-3 py-2 rounded-md">
                    Rules Dashboard
                  </Link>
                  <Link to="/test" className="text-gray-500 hover:text-gray-700 px-3 py-2 rounded-md">
                    Test Console
                  </Link>
                  <Link to="/metrics" className="text-gray-500 hover:text-gray-700 px-3 py-2 rounded-md">
                    Performance
                  </Link>
                </div>
              </div>
            </div>
          </div>
        </nav>

        <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
          <Routes>
            <Route path="/admin" element={<RulesDashboard />} />
            <Route path="/test" element={<TestConsole />} />
            <Route path="/metrics" element={<PerformanceDashboard />} />
            <Route path="/" element={<RulesDashboard />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;