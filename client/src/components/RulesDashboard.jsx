// src/components/RulesDashboard.jsx
import React, { useState, useEffect } from 'react';

const RulesDashboard = () => {
  const [rules, setRules] = useState([]);
  const [newRule, setNewRule] = useState({
    game_id: '',
    countries: [],
    min_version: '',
    platforms: [],
    app_types: [],
    is_active: true
  });

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl mb-4">Game Access Rules Dashboard</h1>
      
      {/* Add New Rule Form */}
      <div className="bg-white p-4 rounded shadow mb-4">
        <h2 className="text-xl mb-2">Add New Rule</h2>
        <form>
          <div className="grid grid-cols-2 gap-4">
            <input 
              placeholder="Game ID"
              className="border p-2 rounded"
            />
            <input 
              placeholder="Countries (comma separated)"
              className="border p-2 rounded"
            />
            <input 
              placeholder="Minimum Version"
              className="border p-2 rounded"
            />
            <select className="border p-2 rounded">
              <option>iOS</option>
              <option>Android</option>
            </select>
          </div>
          <button className="bg-blue-500 text-white px-4 py-2 rounded mt-4">
            Add Rule
          </button>
        </form>
      </div>

      {/* Rules List */}
      <div className="bg-white p-4 rounded shadow">
        <h2 className="text-xl mb-2">Existing Rules</h2>
        <table className="w-full">
          <thead>
            <tr>
              <th>Game ID</th>
              <th>Countries</th>
              <th>Min Version</th>
              <th>Platforms</th>
              <th>Status</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {rules.map(rule => (
              <tr key={rule.id}>
                <td>{rule.game_id}</td>
                <td>{rule.countries.join(', ')}</td>
                <td>{rule.min_version}</td>
                <td>{rule.platforms.join(', ')}</td>
                <td>
                  <span className={`px-2 py-1 rounded ${rule.is_active ? 'bg-green-100' : 'bg-red-100'}`}>
                    {rule.is_active ? 'Active' : 'Inactive'}
                  </span>
                </td>
                <td>
                  <button className="text-blue-500">Edit</button>
                  <button className="text-red-500 ml-2">Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default RulesDashboard;