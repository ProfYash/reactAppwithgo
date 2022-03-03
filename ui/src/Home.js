import React,{useEffect} from 'react'
import AddUser from './AddUser'
import DeleteUser from './DeleteUser'
import DisplayUser from './DisplayUser'
import UpdateUser from './UpdateUser'
import Cookies from 'universal-cookie';
import { useNavigate } from "react-router-dom";

export default ()=>{
   const cook=new Cookies()
    let cookie=cook.get('token')
    let navigate = useNavigate();
    let path='/AllUsers'
    function goTODisplayUser(){
        navigate(path);
    }
    console.log(cookie)
    // return (xml/hml);
    // useEffect(() => {
    //     loadHome();
    // },[])
if (cookie==null){
    return(
        <div>
            <h1>Not Authorized</h1>
        </div>
    )
}else{
     return(
        <div className='container'>
            <div class="row"> 
                <div class="col-sm">
                    <div className="card">
                        <div className="card-body">
                        <h2>Add User from here</h2>
                        <AddUser />
                        </div>
                    </div>
                </div>
                <div class="col-sm">
                    <div className="card">
                        <div className="card-body">
                        <h2>Update User from here</h2>
                        <UpdateUser />
                        </div>
                    </div>
                </div>
                <div class="col-sm">
                    <div className="card">
                        <div className="card-body">
                        <h2>Delete User from here</h2>
                        <DeleteUser />
                        </div>
                    </div>
                </div>
            </div><br /><br />
            <div class="row"> 
                <div class="col-sm">
                    <div className="card">
                        <div className="card-body">     
                            <h2>All Users</h2>
                            <button className="btn btn-primary" onClick={goTODisplayUser}>Click Here</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
     )
 }
}