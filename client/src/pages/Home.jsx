import React from "react"

import Navbar from "../component/navbar"
import Footer from "../component/footer"

function HomePage() {
    return (
        <div>
            <Navbar/>
            <div style={{padding : "20%"}}>
             <h1 className="display-4 mb-4 font-weight-bold text-black">Welcome to Book List Application</h1>
             </div>
            <Footer/>
        </div>
    )
}


export default HomePage