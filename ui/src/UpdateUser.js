import React,{useState} from "react";
import axios from "axios";

export default ()=>{
    // const UID="100"
    const [FName,updateFname]= useState("")
    const [RollNo,updateRollNo]= useState("")
    const [Contact,updateContact]= useState("")
const handleMyUpdate = async(e)=>{
    e.preventDefault();
    if (FName!=""&&RollNo!=""&&Contact!=""){
    await axios.put(`http://localhost:4002/api/v1/blog/UpdateUser/${RollNo}`,{FName,RollNo,Contact}).catch(e=>alert("Roll No not found"))
    updateFname('')
    updateContact('')
    updateRollNo('')}
    else{
        alert("Incomplete Inputs")
    }
}
    



return(
        <form onSubmit={handleMyUpdate}>
        <div>
        <label className="form-group">Roll Number</label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={RollNo} onChange={(e)=>updateRollNo(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">Full Name</label>&nbsp;&nbsp;&nbsp;&nbsp; 
        <input type="text" value={FName} onChange={(e)=>updateFname(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">Mobile</label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={Contact} onChange={(e)=>updateContact(e.target.value)} className="form-control"/><br /><br />
        </div> RollNo:&nbsp;&nbsp;&nbsp;&nbsp;{RollNo}<br/><br/>Fullname:&nbsp;&nbsp;&nbsp;&nbsp;{FName}<br/><br/>Mobile:&nbsp;&nbsp;&nbsp;&nbsp;{Contact}<br/><br/>
        <button className="btn btn-primary">Update</button>
       
    </form>
    )
}