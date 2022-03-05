import React,{useState} from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
export default ()=>{
    // const UID="100"
    let RollNo=useParams().roll
    console.log("RollNo",RollNo)
    let AddName=useParams().add
    console.log("AddName",AddName)
    const [city,updatecity]= useState("")
    const [pincode,updatepincode]= useState("")
    const [firstlineadd,updatefirstlineadd]= useState("")
const handleMyUpdateAddress = async(e)=>{
    e.preventDefault();
   
    if (city!=""&&pincode!=""&&firstlineadd!=""){
    await axios.put(`http://localhost:4002/api/v1/blog/UpdateAddress/${RollNo}/${AddName}`,{city,pincode,firstlineadd}).catch(e=>alert("Roll No not found"))
    updatecity('')
    updatepincode('')
    updatefirstlineadd('')}
    else{
        alert("Incomplete Inputs")
    }
}
    



return(
        <form onSubmit={handleMyUpdateAddress}>
        <div>
        <label className="form-group">Address Line 1:</label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={firstlineadd} onChange={(e)=>updatefirstlineadd(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">City</label>&nbsp;&nbsp;&nbsp;&nbsp; 
        <input type="text" value={city} onChange={(e)=>updatecity(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">Pincode</label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={pincode} onChange={(e)=>updatepincode(e.target.value)} className="form-control"/><br /><br />
        </div> Address Line 1::&nbsp;&nbsp;&nbsp;&nbsp;{firstlineadd}<br/><br/>City:&nbsp;&nbsp;&nbsp;&nbsp;{city}<br/><br/>Pincode:&nbsp;&nbsp;&nbsp;&nbsp;{pincode}<br/><br/>
        <button className="btn btn-primary">Update</button>
       
    </form>
    )
}