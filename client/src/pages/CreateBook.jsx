import React from 'react'

import Navbar from "../component/navbar"
import Footer from "../component/footer"
import { useDispatch } from 'react-redux'
import { useState } from 'react'

import { createBook } from '../store/action/bookAction'

import { useHistory } from "react-router-dom"

function CreateBook() {
    const history = useHistory()
    const dispatch = useDispatch()

    const [title, setTitle] = useState("")
    const [author, setAuthor] = useState("")
    const [year, setYear] = useState(0)

    const submitCreateBook = () => {
        const data = {
            title : title,
            author : author,
            year : year,
        }

        console.log(data)

        dispatch(createBook(data))
    }

    return (
        <div>
            <Navbar/>
            <div className="container">    
            <form style={{textAlign:"center", paddingTop:"150px", paddingBottom:"150px"}} onSubmit={submitCreateBook}>
            <h2>Create New Book</h2>
                <div className="mb-3">
                    <label for="title" className="form-label">Title</label>
                    <input type="text" className="form-control" id="title" onChange={e => {
                        e.preventDefault()
                        setTitle(e.target.value)
                    }}/>
                </div>
                <div className="mb-3">
                    <label for="auhtor" className="form-label">Auhtor</label>
                    <input type="text" className="form-control" id="auhtor" onChange={e => {
                        e.preventDefault()
                        setAuthor(e.target.value)
                    }}/>
                </div>
                <div class="mb-3">
                    <label for="year" className="form-label">Year</label>
                    <input type="text" className="form-control" id="year" onChange={e => {
                        e.preventDefault()
                        setYear(e.target.value)
                    }}/>
                </div>
                <button type="submit" className="btn btn-danger" style={{ margin : "0px 100px"}}>Back</button>
                <button type="submit" className="btn btn-primary" style={{ margin : "0px 100px"}}>Create</button>
            </form>
            </div> 
            <Footer/>
        </div>
    )
}


export default CreateBook
