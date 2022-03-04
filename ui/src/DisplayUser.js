import React,{useState,useEffect} from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import DisplayUserAddress from './DisplayUserAddress'
import AddAddress from "./AddAddress";

export default () => {
    let navigate = useNavigate();
    
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
},[])


const cardofuser = Object.values(users).map(u=>{
    return (
        <div className="card" style={{ width:"30%", marginBottom:"20px"}}>

            <div className="card-body" key={u.UID}>
                Name:&nbsp;&nbsp;&nbsp;{u.FName}<br /><br />
                Roll:&nbsp;&nbsp;&nbsp;{u.RollNo}<br /><br />
                Mobile:&nbsp;&nbsp;&nbsp;{u.Contact}<br /><br />
            </div>
            <div className="card-body" key={u.UID}>
               <DisplayUserAddress roll={u.RollNo} />
            </div>
            <div className="card-body" key={u.UID}>   
            {/* <AddAddress r={u.RollNo}/> */}
                <button className="btn btn-primary" onClick={()=>navigate(`/AddAddress/${u.RollNo}`)}>Add Address</button>
            </div>
        </div>
    )
})

    return(
       
         <div className="d-flex felx-row flex-wrap justify-content-between">
     <h1>{loginStatus}</h1><br />
     {cardofuser}
         </div> 
         )

}