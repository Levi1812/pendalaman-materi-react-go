import React, { useEffect, useState } from "react"
import { useHistory } from 'react-router-dom'
import { useDispatch } from 'react-redux'

import { logoutUser } from '../store/action/userAction'

const Navbar = () => {
    const history = useHistory()

    const dispacth = useDispatch()

    const [token, setToken] = useState("")

    useEffect(() => {
        setToken(localStorage.getItem("access_token"))
    })

    console.log(token)

    const logout = () => {
        dispacth(logoutUser())
        history.push("/login")
    }

    return (
        <header>
            <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
                <div class="container-fluid">
                <a class="navbar-brand" href="#">Book List App</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>

                {!token ? (
                    <div class="collapse navbar-collapse" id="navbarCollapse">
                    <ul class="navbar-nav me-auto mb-2 mb-md-0">
                    <li class="nav-item" onClick={e => history.push("/")}>
                        <a class="nav-link active" aria-current="page">Home</a>
                    </li>
                    <li class="nav-item" onClick={e => history.push("/register")}>
                        <a class="nav-link active" aria-current="page">Register</a>
                    </li>
                    <li class="nav-item" onClick={e => history.push("/login")}>
                        <a class="nav-link active" aria-current="page">Login</a>
                    </li>

                    </ul>
                </div> 
                ) : (
                    <div class="collapse navbar-collapse" id="navbarCollapse">
                    <ul class="navbar-nav me-auto mb-2 mb-md-0">
                    <li class="nav-item" onClick={e => history.push("/dashboard")}>
                        <a class="nav-link active" aria-current="page">Dashboard</a>
                    </li>
                    <li class="nav-item" onClick={logout}>
                        <a class="nav-link active" aria-current="page" href="#">Logout</a>
                    </li>
                    </ul>
                </div>  
                )}
                </div>
            </nav>
        </header>
    )
}


export default Navbar