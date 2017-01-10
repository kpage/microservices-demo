import React, { Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import { fetchOrdersIfNeeded, requestOrdersPage, requestOrderCoffeeForm, cancelUnfinishedOrder } from '../actions'
import OrderList from '../components/OrderList'
import Order from '../components/Order'
//import SVGIcon from 'grommet/components/SVGIcon'
import Java from 'grommet/components/icons/base/Java'
import Title from 'grommet/components/Title'
//import Logo from 'grommet/components/Logo'

class App extends Component {
    constructor(props) {
        super(props)
        this.handlePageClick = this.handlePageClick.bind(this)
        this.handleOrderCoffeeClick = this.handleOrderCoffeeClick.bind(this)
        this.handleCancelUnfinishedOrderClick = this.handleCancelUnfinishedOrderClick.bind(this)
    }

    componentDidMount() {
        const { dispatch } = this.props
        dispatch(fetchOrdersIfNeeded())
    }

    componentWillReceiveProps(nextProps) {
    }

    handlePageClick(linkRef, e) {
        e.preventDefault()

        const { orders, dispatch } = this.props
        dispatch(requestOrdersPage(orders._links[linkRef].href))
    }
    
    handleOrderCoffeeClick(e) {
        e.preventDefault()

        const { dispatch } = this.props
        dispatch(requestOrderCoffeeForm())
    }
    
    handleCancelUnfinishedOrderClick(e) {
        e.preventDefault()

        const { dispatch } = this.props
        dispatch(cancelUnfinishedOrder())
    }

    render() {
        const { orders } = this.props
        const isEmpty = orders.items.length === 0
        return (
            <div>
                {orders.isOrderCoffeeFormOpen
                    ? 
                    <div>
                        <Title onClick={this.handleCancelUnfinishedOrderClick}>Cancel Order</Title>
                    </div>
                    :
                    <div>
                        <Title onClick={this.handleOrderCoffeeClick}>
                            <Java colorIndex="brand" />
                            Order Coffee
                        </Title>
                        {isEmpty
                            ? <h2>Empty.</h2>
                            : <div>
                            <OrderList 
                            items={orders.items}
                            _links={orders._links}
                            page={orders.page}
                            onPageLinkClick={this.handlePageClick}/>
                            </div>
                        }
                    </div>
                }
            </div>
        )
    }
}

App.propTypes = {
    orders: PropTypes.shape({
        _links: PropTypes.object.isRequired, // TODO: more explicit shape
        page: PropTypes.object.isRequired, // TODO: more explicit shape
        items: PropTypes.arrayOf(
            PropTypes.shape(Order.propTypes)
        ).isRequired,
    }),
//  isFetching: PropTypes.bool.isRequired,
//  lastUpdated: PropTypes.number,
    dispatch: PropTypes.func.isRequired
}

function mapStateToProps(state) {
    const { orders } = state
  
    return {
        orders
    }
}

export default connect(mapStateToProps)(App)
