import React, { useState } from 'react';
import SearchForm from './components/SearchForm';
import ProfileCard from './components/ProfileCard';
import PostCard from './components/PostCard';
import { fetchProfiles, fetchPosts } from './api/profiles';

function App() {
  const [profiles, setProfiles] = useState([]);
  const [posts, setPosts] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleSearch = async (companyName, categoryId, keyword) => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await fetchProfiles(companyName, categoryId, keyword);
      setProfiles(data);
    } catch (err) {
      setError('Failed to fetch profiles. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleGetPosts = async (keyword, companyName, categoryId) => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await fetchPosts(keyword, companyName, categoryId);
      setPosts(data);
    } catch (err) {
      setError('Failed to fetch posts. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-100 to-gray-200">
      <div className="container mx-auto p-6">
        <div className="bg-white rounded-lg shadow-lg p-6 mb-8">
          <SearchForm onSearch={handleSearch} onGetPosts={handleGetPosts} />
        </div>
        {isLoading && (
          <div className="flex justify-center items-center">
            <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
          </div>
        )}
        {error && <p className="text-center text-red-500 mb-4">{error}</p>}
        {profiles.length > 0 && (
          <div className="mb-12">
            <h2 className="text-2xl font-bold mb-4 text-gray-800">Profiles</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {profiles.map((profile, index) => (
                <ProfileCard key={profile.Email || index} profile={profile} />
              ))}
            </div>
          </div>
        )}
        {posts.length > 0 && (
          <div>
            <h2 className="text-2xl font-bold mb-4 text-gray-800">Posts</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {posts.map((post, index) => (
                <PostCard key={index} post={post} />
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;