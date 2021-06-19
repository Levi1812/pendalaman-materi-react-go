const initState = {
    books : [],
    book : null,
    isLoading : false,
    error : null
}

const bookReducer = (state = initState, action) => {
    switch(action.type) {
        case "BOOK_LOADING":
            return { ...state, isLoading : true}
        case "FETCH_BOOKS":
            return { ...state, isLoading: false, books : action.payload}
        case "CREATE_BOOK":
            return { ...state, isLoading:false, book : action.payload}
        case "UPDATE_BOOK":
            return { ...state, isLoading:false, book : action.payload}
        case "DELETE_BOOK":
            return { ...state, isLoading:false, book : null}
        case "ERROR_BOOKS":
            return { ...state, isLoading: false}
        default:
            return state
    }
}

export default bookReducer