import React from 'react';
import './App.css';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Home from "./pages/Home";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={
          <Home />
        }/>
        <Route path='/hello' element={
          <div>hello</div>
        }/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
