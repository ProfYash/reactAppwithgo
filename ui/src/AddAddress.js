import React,{useState} from "react";
import axios from "axios";

export default()=> {
const UID="100"
const addid="1000"
const [addressname,updateaddressname]=useState('')
const [firstlineadd,updatefirstline]=useState('')
const [city,updatecity]=useState('')
const [pincode,updatepincode]=useState('')
const [RollNo,updateRollNo]= useState('')
const handleMyAddress=async(e)=>{
    if (firstlineadd!=""&&RollNo!=""&&city!=""&&pincode!=""){

       
        e.preventDefault();
        await axios.post(`http://localhost:4002/api/v1/blog/addaddress/${RollNo}`,{addid,UID,addressname,firstlineadd,city,pincode}).catch(e=>console.log(e.message))
        updatecity('')
        updatepincode('')
        updateRollNo('')
        updatefirstline('')
    }
        else{
            alert("Incomplete Inputs")
        }
}


    return(
        <form onSubmit={handleMyAddress}>
            <div>
        <label className="form-group">Roll Number</label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={RollNo} onChange={(e)=>updateRollNo(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">AddressName:</label>&nbsp;&nbsp;&nbsp;&nbsp; 
        <input type="text" value={addressname} onChange={(e)=>updateaddressname(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">Address Line 1:</label>&nbsp;&nbsp;&nbsp;&nbsp; 
        <input type="text" value={firstlineadd} onChange={(e)=>updatefirstline(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">City:</label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={city} onChange={(e)=>updatecity(e.target.value)} className="form-control"/><br /><br />
        <label className="form-group">Pincode:</label>&nbsp;&nbsp;&nbsp;&nbsp;
        <input type="text" value={pincode} onChange={(e)=>updatepincode(e.target.value)} className="form-control"/><br /><br />
        </div> RollNo:&nbsp;&nbsp;&nbsp;&nbsp;{RollNo}
            
            <button className="btn btn-primary">Add User</button>
           
        </form>
    )
}