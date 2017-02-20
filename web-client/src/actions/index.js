import fetch from 'isomorphic-fetch'
import client from '../client'
import follow from '../follow' // function to hop multiple links by "rel"
import post from '../post' // function to hop multiple links by "rel"
const root = '/api';
// Can remove this separate api root once server implements a top-level HATEOAS root that includes links to /api and /auth
const authRoot = '/auth';

export const REQUEST_ORDERS = 'REQUEST_ORDERS'
export const RECEIVE_ORDERS = 'RECEIVE_ORDERS'
export const REQUEST_ORDERS_PAGE = 'REQUEST_ORDERS_PAGE'
export const REQUEST_ORDER_COFFEE_FORM = 'REQUEST_ORDER_COFFEE_FORM'
export const CANCEL_UNFINISHED_ORDER = 'CANCEL_UNFINISHED_ORDER'
export const LOGGING_IN = 'LOGGING_IN'
export const LOGGED_IN = 'LOGGED_IN'
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
        return follow(client, root, [{rel: 'restbucks:orders', params: {size: 5}}])
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

export function loggingIn() {
    return {
        type: LOGGING_IN
    }
}

export function loggedIn(token) {
    return {
        type: LOGGED_IN,
        "token": token
    }
}

export function login() {
    return dispatch => {
        dispatch(loggingIn())
        return post(client, authRoot, [{rel: 'auth:token', entity: {username: "testuser", password: "testpassword"}}])
            .then(response => dispatch(loggedIn(response)))
    }
}