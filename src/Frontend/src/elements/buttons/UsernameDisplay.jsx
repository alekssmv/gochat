import React from 'react';

const UsernameDisplay = () => {

  const handleUsernameDisplay = async () => {

    // Define username
    let username;

    // Get the username from /username
    await fetch('/user', {
      method: 'POST',
      credentials: 'include'
    }).then(res => res.json()).then(data => username = data.username);


  };

  return (
    <p>Username {handleUsernameDisplay}</p>
  );
};

export default Logout;
