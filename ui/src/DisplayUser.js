import React,{useState,useEffect} from "react";
import axios from "axios";
export default () => {
const [loginStatus,updateloginStatus]=useState('')
    const [users,updateusers] = useState({}) // creating a hook for updates
const loadusers = async ()=>{
   const resp = await axios.get("http://localhost:4002/api/v1/blog/getUser").catch(e=>{
       console.log(e.message)
       if (e.response.status==401){
        updateloginStatus('Unauthorized')        
        return
    } 

})

if( resp!=null){
   console.log("data: ",resp.data)
   updateusers(resp.data)
}
}
useEffect(() => {
    loadusers();
},[])//useeffect for call method at time of load and only once

//creating dynmic number of cards
const cardofuser = Object.values(users).map(u=>{
    return (
        <div className="card" style={{ width:"30%", marginBottom:"20px"}}>

<div className="card-body" key={u.UID}>
{/* ID:&nbsp;&nbsp;&nbsp;{u.UID}<br /><br /> */}
    Name:&nbsp;&nbsp;&nbsp;{u.FName}<br /><br />
    Roll:&nbsp;&nbsp;&nbsp;{u.RollNo}<br /><br />
    Mobile:&nbsp;&nbsp;&nbsp;{u.Contact}<br /><br />
{/* hello world */}
</div>
        </div>
    )
})
// if (users!={}){
    return(
        // <div><h1>Display</h1></div> //do this initially
         <div className="d-flex felx-row flex-wrap justify-content-between">
     <h1>{loginStatus}</h1><br />
     {cardofuser}
         </div> 
         )

}