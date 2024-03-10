import axios from 'axios';
import React from 'react';

function LogoutPage() {
  async function handleLogout() {
    try {
      const response = await axios.get('http://localhost:3000/logout/google', {
        withCredentials: true,
      });
      console.log('response -' + response.data);
    } catch (error) {
      console.log('Error - ' + error);
    }
  }
  return (
    <div>
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
}

export default LogoutPage;
