import React, { useState } from 'react';

const categories = [
  { id: 12, name: 'Human Resources personnel' },
  { id: 8, name: 'Engineers' },
  { id: 18, name: 'Program Managers' },
  { id: 24, name: 'Researchers' },
  { id: 13, name: 'IT Personnel' },
];

function SearchForm({ onSearch }) {
  const [companyName, setCompanyName] = useState('');
  const [categoryId, setCategoryId] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    if (companyName && categoryId) {
      onSearch(companyName, parseInt(categoryId, 10));
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mb-4">
      <div className="flex flex-col md:flex-row gap-2">
        <input
          type="text"
          value={companyName}
          onChange={(e) => setCompanyName(e.target.value)}
          placeholder="Enter company name"
          className="border p-2 rounded"
          required
        />
        <select
          value={categoryId}
          onChange={(e) => setCategoryId(e.target.value)}
          className="border p-2 rounded"
          required
        >
          <option value="">Select a category</option>
          {categories.map((category) => (
            <option key={category.id} value={category.id}>
              {category.name}
            </option>
          ))}
        </select>
        <button type="submit" className="bg-blue-500 text-white p-2 rounded">
          Search
        </button>
        <button 
          className="mt-4 bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded transition duration-300 ease-in-out"
        >
          Get Posts
        </button>
      </div>
      
    </form>
  );
}

export default SearchForm;
