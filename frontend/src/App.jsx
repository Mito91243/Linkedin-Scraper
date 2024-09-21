import React, { useState } from 'react';
import SearchForm from './components/SearchForm';
import ProfileCard from './components/ProfileCard';
import { fetchProfiles } from './api/profiles';

function App() {
  const [profiles, setProfiles] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleSearch = async (companyName, categoryId) => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await fetchProfiles(companyName, categoryId);
      setProfiles(data);
    } catch (err) {
      setError('Failed to fetch profiles. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleGetPosts = () => {
    // Implement the logic to fetch posts here
    console.log('Fetching posts...');
  };

  return (
    <div className="container mx-auto p-6 bg-gray-100 min-h-screen">
      <div className="bg-white rounded-lg shadow-md p-6 mb-6">
        <SearchForm onSearch={handleSearch} />

      </div>
      {isLoading && <p className="text-center text-gray-600">Loading...</p>}
      {error && <p className="text-center text-red-500 mb-4">{error}</p>}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {profiles.map((profile, index) => (
          <ProfileCard 
            key={profile.Email || index} 
            profile={profile}
          />
        ))}
      </div>
    </div>
  );
}

export default App;