import React from 'react'
import { Route,Link, Routes } from 'react-router-dom'
import AddAddress from './AddAddress'
import DisplayUser from './DisplayUser'
import DisplayUserAddress from './DisplayUserAddress'
import Home from './Home'
import Login from './Login'

function App(){
    return(
        <div>
            <Routes>
            <Route exact path="/home" component={Home} element={<Home />} />
            <Route exact path="/" component={Login} element={<Login />} />
            <Route exact path="/AllUsers" component={DisplayUser} element={<DisplayUser />} />
            <Route exact path="/AddAddress/:roll" component={AddAddress} element={<AddAddress />} />
            <Route exact path="/DisplayAddress/:roll" component={DisplayUserAddress} element={<DisplayUserAddress />} />
            <Route path="*" element={<Login />} />
            </Routes>
        </div>
    )
}
export default App