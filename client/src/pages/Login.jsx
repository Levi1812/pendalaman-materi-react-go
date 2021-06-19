import React, {useState} from "react"

import Navbar from "../component/navbar"
import Footer from "../component/footer"

import { useDispatch, useSelector } from 'react-redux'

import { loginUser } from '../store/action/userAction'
import { useHistory } from "react-router-dom"

function LoginPage() {

    const history = useHistory()

    const dispatch = useDispatch()
    const [email, setEmail] = useState("")
    const [pass, setPass] = useState("")

    const { error } = useSelector(state => state.user)

    const loginSubmit = () => {
        const data = {
            email : email,
            password :pass,
        }

        console.log(data)

        if (!error) {
            dispatch(loginUser(data))
            history.push("/dashboard")
        }

    }


    return ( 
        <div>
            <Navbar/>
            <div className="container">    
            <form style={{textAlign:"center", paddingTop:"150px", paddingBottom:"150px"}} onSubmit={loginSubmit}>
            <h2>Login User</h2>
                <div className="mb-3">
                    <label for="email" className="form-label">Email address</label>
                    <input type="email" className="form-control" id="email" aria-describedby="emailHelp"onChange={e => {
                    e.preventDefault()
                    setEmail(e.target.value)
                }}/>
                    <div id="emailHelp" className="form-text">We'll never share your email with anyone else.</div>
                </div>
                <div class="mb-3">
                    <label for="password" className="form-label">Password</label>
                    <input type="password" className="form-control" id="password" onChange={e => {
                    e.preventDefault()
                    setPass(e.target.value)
                }}/>
                </div>
                <button type="submit" className="btn btn-primary">Login</button>
            </form>
            </div> 
            <Footer/>
        </div>
    )
}


export default LoginPage