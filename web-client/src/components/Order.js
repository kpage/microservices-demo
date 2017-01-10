import React, { PropTypes } from 'react'
import moment from 'moment'

const Order = ({ customerName, orderedDate, status }) => (
    <tr>
        <td>{customerName}</td>
        <td>{status}</td>
        <td>{ moment(orderedDate).fromNow() }</td>
    </tr>
)

Order.propTypes = {
    customerName: PropTypes.string.isRequired,
    orderedDate: PropTypes.string.isRequired,
    status: PropTypes.string.isRequired
//  onClick: PropTypes.func.isRequired,
//  completed: PropTypes.bool.isRequired,
//  text: PropTypes.string.isRequired
}

export default Order
