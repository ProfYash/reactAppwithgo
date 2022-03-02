import React from 'react'
import AddUser from './AddUser'
import DeleteUser from './DeleteUser'
import DisplayUser from './DisplayUser'
import UpdateUser from './UpdateUser'
export default ()=>{
    // return (xml/hml);
 
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
                            <DisplayUser />
                        </div>
                    </div>
                </div>
            </div>
        </div>
     )
 }