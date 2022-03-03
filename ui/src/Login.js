import React,{useState} from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useCookies} from 'react-cookie'
export default()=>{
    const [username,updateUsername]=useState('')
    const [loginStatus,updateloginStatus]=useState('')
const [password,updatepassword]=useState('')
let navigate = useNavigate(); 
let path = `/home`; 
const [cookies, setCookie] = useCookies('token');
const handleMyLogin= async(e) =>{

    if (username!=""&&password!=""){
        e.preventDefault();
        let resp=await axios.post(`http://localhost:4002/login`,{username,password}).catch(e=>{
            console.log(e.message)
            if (e.response.status==401){
                updateloginStatus("login failed")
                return
            } 
    })
    if( resp!=null){
    updateloginStatus("")
                console.log("Here")
                navigate(path);
      
        updateUsername('')
    }

    }
}


    return(
        <form onSubmit={handleMyLogin}>
            <div>
        <label className="form-group">Username: </label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={username} onChange={(e)=>updateUsername(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">Password</label>&nbsp;&nbsp;&nbsp;&nbsp; 
        <input type="text" value={password} onChange={(e)=>updatepassword(e.target.value)} className="form-control"/><br /><br />
        
    </div>
    <button className="btn btn-primary">Login Here</button>
    {/* <button className="btn btn-primary" onClick={goToHome}>goToHome</button> */}
    <div>
    Username:&nbsp;&nbsp;&nbsp;&nbsp;{username}<br />{loginStatus}
    </div>
    </form>
    )

}