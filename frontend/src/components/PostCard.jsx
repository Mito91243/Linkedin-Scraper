import React from 'react';

function PostCard({ post }) {
  return (
    <div className="bg-white rounded-lg shadow-lg overflow-hidden">
      <div className="bg-gradient-to-r from-blue-500 to-purple-600 p-4">
        <h3 className="text-xl font-semibold text-white truncate">{post.name || 'Anonymous'}</h3>
      </div>
      <div className="p-6">
        <p className="text-gray-700 mb-4">{post.text}</p>
        <div className="flex items-center justify-between text-sm text-gray-500">
          <div className="flex items-center space-x-2">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
              <path fillRule="evenodd" d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z" clipRule="evenodd" />
            </svg>
            <span>{post.numLikes || 0}</span>
          </div>
          <div className="flex items-center space-x-2">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
              <path fillRule="evenodd" d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z" clipRule="evenodd" />
            </svg>
            <span>{post.numComments || 0}</span>
          </div>
        </div>
        {post.date && (
          <p className="text-sm text-gray-500 mt-4">Posted: {post.date}</p>
        )}
        {/* Hiding the URN field for now */}
        {/* {post.urn && (
          <p className="text-sm text-gray-500 mt-2">URN: {post.urn}</p>
        )} */}
      </div>
      {post.actionTarget && (
        <a
          href={post.actionTarget}
          target="_blank"
          rel="noopener noreferrer"
          className="block bg-blue-500 text-white text-center py-2 hover:bg-blue-600 transition duration-300"
        >
          Apply for this Job
        </a>
      )}
    </div>
  );
}

export default PostCard;
