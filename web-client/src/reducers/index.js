import { combineReducers } from 'redux'
import * as actions from '../actions'

// The auth reducer. The starting state sets authentication
// based on a token being in local storage. In a real app,
// we would also want a util to check if the token is expired.
function auth(state = {
    isFetching: false,
    // TODO: is it cool to pull from localStorage in a reducer?  Should this go
    // in the state instead or would that allow browser and redux store state to go out of sync?
    isAuthenticated: localStorage.getItem('access_token') ? true : false
}, action) {
    switch (action.type) {
        case actions.LOGIN_REQUEST:
            return Object.assign({}, state, {
                isFetching: true,
                isAuthenticated: false,
                user: action.creds
            })
        case actions.LOGIN_SUCCESS:
            return Object.assign({}, state, {
                isFetching: false,
                isAuthenticated: true,
                errorMessage: ''
            })
        case actions.LOGIN_FAILURE:
            return Object.assign({}, state, {
                isFetching: false,
                isAuthenticated: false,
                errorMessage: action.message
            })
        case actions.LOGOUT_SUCCESS:
            return Object.assign({}, state, {
                isFetching: true,
                isAuthenticated: false
            })
        default:
            return state
    }
}

function orders(state = {
    //  isFetching: false,
    //  didInvalidate: false,
    items: [],
    _links: {},
    page: {},
    isOrderCoffeeFormOpen: false,
}, action) {
    switch (action.type) {
        //    case INVALIDATE_REDDIT:
        //      return Object.assign({}, state, {
        //        didInvalidate: true
        //      })
        case actions.REQUEST_ORDERS:
            return Object.assign({}, state, {
                //        isFetching: true,
                //        didInvalidate: false
            })
        case actions.REQUEST_ORDERS_PAGE:
            return Object.assign({}, state, {
                //          isFetching: true,
                //          didInvalidate: false
            })
        case actions.RECEIVE_ORDERS:
            return Object.assign({}, state, {
                //        isFetching: false,
                //        didInvalidate: false,
                //           TODO: is there ES6 shortcut syntax for copying same props with same names?
                items: action.items,
                _links: action._links,
                page: action.page,
                //        lastUpdated: action.receivedAt
            })
        case actions.REQUEST_ORDER_COFFEE_FORM:
            return Object.assign({}, state, {
                isOrderCoffeeFormOpen: true
            })
        case actions.CANCEL_UNFINISHED_ORDER:
            return Object.assign({}, state, {
                isOrderCoffeeFormOpen: false
            })
        default:
            return state
    }
}

const rootReducer = combineReducers({
    auth,
    orders
})

export default rootReducer