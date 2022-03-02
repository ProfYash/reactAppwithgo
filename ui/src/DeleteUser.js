import React,{useState} from "react";
import axios from "axios";

export default ()=>{

const [RollNo,updateRollNo]=useState("")
// const resp
const handleDelete = async(e)=>{
    e.preventDefault();
    if (RollNo!=""){
  let  resp=await axios.delete(`http://localhost:4002/api/v1/blog/deleteuser/${RollNo}`,{RollNo})
  .catch(e=>alert("Roll No Not Found"))
//   alert(resp.status)
    // console.log("response:",resp)
    updateRollNo("")}
    else{
        alert("Incomplete Inputs")
    }
}
return (

    <form onSubmit={handleDelete}>
        <div>
            <label className="form-group">Roll Number</label>&nbsp;&nbsp;&nbsp;&nbsp; 
            <input type="text" value={RollNo} onChange={(e)=>updateRollNo(e.target.value)} className="form-control"/><br /><br />
            <label className="form-group">RollNo you want to delete : {RollNo}</label><br /><br />
            <button className="btn btn-primary">Delete</button>
            </div>
    </form>
)



}