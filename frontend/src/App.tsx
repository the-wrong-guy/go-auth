import './App.css';
import { Routes, Route } from 'react-router-dom';
import HomePage from './pages/home';
import LoginPage from './pages/login';
import LogoutPage from './pages/logout';

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/logout" element={<LogoutPage />} />
    </Routes>
  );
}

export default App;
