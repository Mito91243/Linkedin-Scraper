import React from 'react';

function ProfileCard({ profile }) {
  const isAnonymous = profile.fullName === "LinkedIn Member";

  return (
    <div className="border rounded p-4 shadow">
      <h2 className="text-xl font-semibold mb-2">
        {isAnonymous ? 'Anon' : profile.fullName}
      </h2>
      <p><strong>Position:</strong> {profile.position || '--'}</p>
      <p><strong>Email:</strong> {isAnonymous ? 'N/A' : (profile.Email || 'N/A')}</p>
      <a
        href={profile.Link}
        target="_blank"
        rel="noopener noreferrer"
        className="text-blue-500 hover:underline mt-2 inline-block"
      >
        View Full Profile
      </a>
    </div>
  );
}

export default ProfileCard;