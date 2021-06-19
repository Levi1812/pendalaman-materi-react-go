// semua action untuk book ada disini

import apiBase from '../../API/axios'

const access_token = !localStorage.getItem("access_token") ? "" : localStorage.getItem("access_token")

export const fetchBooks = () => {
    return async (dispatch) => {
        try {
            dispatch({type : "BOOK_LOADING"})

            const { data } = await apiBase({
                method : "GET",
                url : "/books",
                headers : {
                    "Authorization" : access_token
                }
            })

            console.log(data)

           return dispatch({ type : "FETCH_BOOKS", payload : data})

        } catch(err) {
            dispatch({ type : "ERROR_BOOKS"})
            console.log(err.response.data)
        }
    }
}


export const createBook = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({type : "BOOK_LOADING"})

            const { data } =  await apiBase({
                method : "POST",
                url : "/books",
                data : payload,
                headers : {
                    "Authorization" : access_token
                }
            })

            console.log(data)

            return dispatch({ type : "CREATE_BOOK", payload : data})

        } catch(err) {
            dispatch({ type : "ERROR_BOOKS"})
            console.log(err.response.data)
        }
    }

}


export const updateBook = (id, payload) => {
    return async (dispatch) => {
        try {
            dispatch({type : "BOOK_LOADING"})

            const { data } = await apiBase({
                method : "PUT",
                url : `/books/${id}`,
                data : payload,
                headers : {
                    "Authorization" : access_token
                }
            })

            console.log(data)

            return dispatch({ type : "UPDATE_BOOK", payload : data})

        } catch(err) {
            dispatch({ type : "ERROR_BOOKS"})
            console.log(err.response.data)
        }
    }
}


export const deleteBook = (id) => {
    return async (dispatch) => {
        try {
            dispatch({type : "BOOK_LOADING"})

            const { data } =  await apiBase({
                method : "DELETE",
                url : `/books/${id}`,
                headers : {
                    "Authorization" : access_token
                }
            })

            console.log(data)

           return dispatch({ type : "DELETE_BOOK", payload : data})
        } catch(err) {
            dispatch({ type : "ERROR_BOOKS"})
            console.log(err.response.data)
        }
    }

}