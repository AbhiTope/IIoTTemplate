import './App.css'
import Home from './components/Home'
import Login from './components/Login'
import AppLayout from './components/AppLayout'

import { BrowserRouter, Routes, Route } from "react-router-dom";

function App() {
  return (
      <>
     <BrowserRouter>
      <Routes>
        <Route path="/" element={<AppLayout />}>
          <Route index element={<Home />} />
        </Route>
      <Route path="/login" element={<Login />} />
      </Routes>
    </BrowserRouter>      </>
  )
}

export default App
