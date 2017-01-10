import { combineReducers } from 'redux'
import * as actions from '../actions'

//function selectedReddit(state = 'reactjs', action) {
//  switch (action.type) {
//    case SELECT_REDDIT:
//      return action.reddit
//    default:
//      return state
//  }
//}
//
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

//function postsByReddit(state = { }, action) {
//  switch (action.type) {
////    case INVALIDATE_REDDIT:
//    case RECEIVE_ORDERS:
//    case REQUEST_ORDERS:
//      return Object.assign({}, state, {
//        [action.reddit]: posts(state[action.reddit], action)
//      })
//    default:
//      return state
//  }
//}

const rootReducer = combineReducers({
//  postsByReddit,
//  selectedReddit
	orders
})

export default rootReducer