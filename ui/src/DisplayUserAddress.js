import React,{useState,useEffect} from "react";
import axios from "axios";
import DisplayUserAddress from './DisplayUserAddress'
import {useParams} from "react-router-dom"
import { useNavigate } from "react-router-dom";
import DeleteUserAdd from "./DeleteUserAdd";
export default ({roll}) => {
    let navigate = useNavigate();
const [loginStatus,updateloginStatus]=useState('')
    const [address,updateaddress] = useState({})
    let RollNo=roll 
console.log(RollNo)
   
const loadusers = async (e)=>{
    
   const resp = await axios.get(`http://localhost:4002/api/v1/blog/addaddress/${RollNo}`).catch(e=>{
       console.log(e.message)
       if (e.response.status==401){
        updateloginStatus('Unauthorized')        
        return
    } 

})

if( resp!=null){
   console.log("data: ",resp.data)
   updateaddress(resp.data)
}
}
useEffect(() => {
    
    loadusers();
},[])//useeffect for call method at time of load and only once

//creating dynmic number of cards
const cardofuser = Object.values(address).map(a=>{
    return (
        <div className="card" style={{ width:"50%", marginBottom:"50px"}}>

<div className="card-body" key={a.addid}>
{/* ID:&nbsp;&nbsp;&nbsp;{u.UID}<br /><br /> */}
    AddressName:&nbsp;&nbsp;&nbsp;{a.addressname}<br /><br />
    First Line:&nbsp;&nbsp;&nbsp;{a.firstlineadd}<br /><br />
    City:&nbsp;&nbsp;&nbsp;{a.city}<br /><br />
    Pincode:&nbsp;&nbsp;&nbsp;{a.pincode}<br /><br />
    <div className="card-body" key={a.UID}>   
            {/* <AddAddress r={u.RollNo}/> */}
                {/* <button className="btn btn-primary" onClick={()=>navigate(`/DeleteAddress/${RollNo}/${a.addressname}`)}>Delete Address</button> */}
            
            <DeleteUserAdd roll={RollNo} add={a.addressname} />
            <button className="btn btn-primary" onClick={()=>navigate(`/UpdateAddress/${RollNo}/${a.addressname}`)}>Update Address</button>
            </div>
    
{/* hello world */}
</div>
        </div>
    )
})
// if (users!={}){
    return(
        // <div><h1>Display</h1></div> //do this initially
         <div className="d-flex felx-row flex-wrap justify-content-between">

     {cardofuser}
         </div> 
         )

}