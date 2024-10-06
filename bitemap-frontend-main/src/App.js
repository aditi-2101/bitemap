
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { useEffect, useState } from 'react';
import ReviewPage from './ReviewPage';
import SignUp from './Signup';
import SignIn from './Signin';
import Map from './Map';
import Home from './Home';


function App() {

  const [loggedIn, setLoggedIn] = useState(false)
  const [email, setEmail] = useState("")


  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<SignIn/>} />
          <Route path="/home" element={<Home/>} />
          <Route path="/signup" element={ <SignUp/>} />
          <Route path="/map" element={ <Map/>} />
          <Route path="/review" element={<ReviewPage />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
