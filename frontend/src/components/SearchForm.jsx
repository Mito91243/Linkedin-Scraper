import React, { useState } from 'react';

const categories = [
  { id: 12, name: 'Human Resources personnel' },
  { id: 8, name: 'Engineers' },
  { id: 18, name: 'Program Managers' },
  { id: 24, name: 'Researchers' },
  { id: 13, name: 'IT Personnel' },
];

function SearchForm({ onSearch, onGetPosts }) {
  const [companyName, setCompanyName] = useState('');
  const [categoryId, setCategoryId] = useState('');
  const [keyword, setKeyword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    if (companyName && categoryId) {
      onSearch(companyName, parseInt(categoryId, 10), keyword);
    }
  };

  const handleGetPosts = () => {
    if (companyName && categoryId && keyword) {
      onGetPosts(keyword, companyName, parseInt(categoryId, 10));
    } else {
      alert("Please fill in all fields (company, category, and keyword) before getting posts.");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mb-4">
      <div className="flex flex-col md:flex-row gap-2 mb-4">
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
        <input
          type="text"
          value={keyword}
          onChange={(e) => setKeyword(e.target.value)}
          placeholder="Enter keyword"
          className="border p-2 rounded"
        />
      </div>
      <div className="flex flex-col md:flex-row gap-2">
        <button type="submit" className="bg-blue-500 text-white p-2 rounded">
          Search Profiles
        </button>
        <button 
          type="button"
          onClick={handleGetPosts}
          className="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded transition duration-300 ease-in-out"
        >
          Get Posts
        </button>
      </div>
    </form>
  );
}

export default SearchForm;