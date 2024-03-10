function LoginPage() {
  function handleLogin() {
    window.location.href = 'http://localhost:3000/auth/google';
  }
  return (
    <div>
      <button onClick={handleLogin}>Login</button>
    </div>
  );
}

export default LoginPage;
