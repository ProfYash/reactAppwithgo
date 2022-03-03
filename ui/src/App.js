import React from 'react'
import { Route,Link, Routes } from 'react-router-dom'
import DisplayUser from './DisplayUser'
import Home from './Home'
import Login from './Login'

function App(){
    return(
        <div>
            <Routes>
            <Route exact path="/home" component={Home} element={<Home />} />
            <Route exact path="/" component={Login} element={<Login />} />
            <Route exact path="/AllUsers" component={DisplayUser} element={<DisplayUser />} />
            
            </Routes>
        </div>
    )
}
export default App