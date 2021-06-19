import React, { useEffect } from "react"

import Navbar from "../component/navbar"
import Footer from "../component/footer"

import { fetchBooks } from '../store/action/bookAction'
import { useSelector, useDispatch } from "react-redux"
import { useHistory } from "react-router-dom"

function Dashboard() {

    const history = useHistory()
    const dispatch = useDispatch()

    const { books, isLoading } = useSelector(state => state.book)

    useEffect(() => {
        dispatch(fetchBooks())
    },[])

    // console.log(isLoading)
    console.log(books)

    // if (isLoading){
    //     return (<h1>Loading</h1>) 
    // } else {

    const gotoCreate = () => {
        history.push("/books")
    }

        return(
            <div>
                <Navbar/>
                <div style={{padding: "10%"}}>
                    <button class="btn btn-primary" onClick={gotoCreate}>Create New Book</button>
                    <table class="table table-hover">
                        <thead>
                            <tr>
                            <th scope="col">First</th>
                            <th scope="col">Last</th>
                            <th scope="col">Handle</th>
                            <th scope="col">Update</th>
                            <th scope="col">Delete</th>
                            </tr>
                        </thead>
                        <tbody>
                            { books && books.map((book, idx) => {
                                return (
                                    <tr>
                                        <td>{book.title}</td>
                                        <td>{book.author}</td>
                                        <td>{book.year}</td>
                                        <td>
                                        <button class="btn btn-warning">Update</button>
                                        </td>
                                        <td>
                                        <button class="btn btn-danger">Delete</button>
                                        </td>
                                    </tr>
                                )
                            })} 
                        </tbody>
                    </table>
                </div>
            <Footer/>
            </div>
        )
    // }
}


export default Dashboard