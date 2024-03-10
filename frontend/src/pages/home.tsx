import React from 'react';
import axios from 'axios';

function HomePage() {
  async function callCheckApi() {
    try {
      const response = await axios.get('http://localhost:3000/check', {
        withCredentials: true,
      });
      console.log('response -' + response);
    } catch (error) {
      console.log('Error - ' + error);
    }
  }

  return (
    <div>
      <h3>Home</h3>
      <button onClick={callCheckApi}>Check API</button>
    </div>
  );
}

export default HomePage;
