import React, { useState, useEffect } from 'react';

const UsernameDisplay = () => {
  // State to hold the username
  const [username, setUsername] = useState('');

  // Effect to fetch the username once the component is mounted
  useEffect(() => {
    const fetchUsername = async () => {
      try {
        console.log('Fetching username...');
        const response = await fetch('/username', {
          method: 'GET',
          credentials: 'include',
        });
        const data = await response.json();
        setUsername(data.username);
      } catch (error) {
        console.error('Error fetching username:', error);
      }
    };

    fetchUsername();
  }, []); // Empty dependency array ensures this effect runs only once

  return <p>Username: {username}</p>;
};

export default UsernameDisplay;
