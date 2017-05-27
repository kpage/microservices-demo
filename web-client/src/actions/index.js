import fetch from 'isomorphic-fetch'
import client from '../client'
import follow from '../follow' // function to hop multiple links by "rel"
const root = '/api';

export const REQUEST_ORDERS = 'REQUEST_ORDERS'
export const RECEIVE_ORDERS = 'RECEIVE_ORDERS'
export const REQUEST_ORDERS_PAGE = 'REQUEST_ORDERS_PAGE'
export const REQUEST_ORDER_COFFEE_FORM = 'REQUEST_ORDER_COFFEE_FORM'
export const CANCEL_UNFINISHED_ORDER = 'CANCEL_UNFINISHED_ORDER'
export const LOGIN_REQUEST = 'LOGIN_REQUEST'
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS'
export const LOGIN_FAILURE = 'LOGIN_FAILURE'
export const LOGOUT_REQUEST = 'LOGOUT_REQUEST'
export const LOGOUT_SUCCESS = 'LOGOUT_SUCCESS'
export const LOGOUT_FAILURE = 'LOGOUT_FAILURE'
//export const SELECT_REDDIT = 'SELECT_REDDIT'
//export const INVALIDATE_REDDIT = 'INVALIDATE_REDDIT'

//export function selectReddit(reddit) {
//  return {
//    type: SELECT_REDDIT,
//    reddit
//  }
//}

//export function invalidateReddit(reddit) {
//  return {
//    type: INVALIDATE_REDDIT,
//    reddit
//  }
//}

function requestOrders() {
    return {
        type: REQUEST_ORDERS
    }
}

export function requestOrdersPage(href) {
    return dispatch => {
        dispatch({ type: REQUEST_ORDERS_PAGE })
        return client({method: 'GET', path: href}).done(response => {
            dispatch(receiveOrders(response.entity))
        });
    }
}

function receiveOrders(orders) {
    return {
        type: RECEIVE_ORDERS,
        items: orders._embedded['restbucks:orders'],
        _links: orders._links,
        page: orders.page,
        receivedAt: Date.now()
    }
}

function fetchOrders() {
    // redux thunk middleware allows us to return a function that is called with 'dispatch', which is just a convenience so 
    // that we don't have to pass 'dispatch' ourselves
    return dispatch => {
        dispatch(requestOrders())
        // hard-code page size of 5 for now, but this could be made into a parameter
        // TODO: have 'follow' pass access_token for every request, not individually
        return follow(client, root, [{rel: 'restbucks:orders', params: {size: 5, access_token: localStorage.getItem('access_token')}}])
            .then(response => response.entity)
            .then(json => dispatch(receiveOrders(json)))
    }
}

// TODO: re-implement the "if needed" part, fetch the correct page number
export function fetchOrdersIfNeeded() {
    return (dispatch) => {
        return dispatch(fetchOrders())
    }
}

export function requestOrderCoffeeForm() {
    return {
        type: REQUEST_ORDER_COFFEE_FORM
    }
}

export function cancelUnfinishedOrder() {
    return {
        type: CANCEL_UNFINISHED_ORDER
    }
}

function requestLogin(creds) {
    return {
        type: LOGIN_REQUEST,
        isFetching: true,
        isAuthenticated: false,
        creds
    }
}

function receiveLogin(user) {
    return {
        type: LOGIN_SUCCESS,
        isFetching: false,
        isAuthenticated: true,
        id_token: user.id_token
    }
}

function loginError(message) {
    return {
        type: LOGIN_FAILURE,
        isFetching: false,
        isAuthenticated: false,
        message
    }
}

// Calls the API to get a token and
// dispatches actions along the way
export function loginUser(creds) {
    // TODO: check that these are the correct headers and body to send to our auth endpoint
    let config = {
        method: 'POST',
        headers: { 'Content-Type':'application/x-www-form-urlencoded' },
        body: `username=${creds.username}&password=${creds.password}`
    }

  return dispatch => {
    // We dispatch requestLogin to kickoff the call to the API
    dispatch(requestLogin(creds))

    // TODO: remove hard-coded absolute url with localhost
    return fetch('http://localhost/auth/token', config)
            .then(response =>
                // TODO: is "user" the best name for this response entity?  maybe "token"?
                response.json().then(user => ({ user, response }))
            ).then(({ user, response }) =>  {
                if (!response.ok) {
                    // If there was a problem, we want to
                    // dispatch the error condition
                    dispatch(loginError(user.message))
                    return Promise.reject(user)
                } else {
                    // If login was successful, set the token in local storage
                    localStorage.setItem('access_token', user.access_token)
                    // Dispatch the success action
                    dispatch(receiveLogin(user))
                }
            }).catch(err => console.log("Error: ", err))
    }
}

function requestLogout() {
  return {
    type: LOGOUT_REQUEST,
    isFetching: true,
    isAuthenticated: true
  }
}

function receiveLogout() {
  return {
    type: LOGOUT_SUCCESS,
    isFetching: false,
    isAuthenticated: false
  }
}

export function logoutUser() {
  return dispatch => {
    dispatch(requestLogout())
    localStorage.removeItem('access_token')
    dispatch(receiveLogout())
  }
}