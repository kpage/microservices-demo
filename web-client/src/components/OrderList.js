import React, { PropTypes } from 'react'
import Anchor from 'grommet/components/Anchor'
import ChapterNext from 'grommet/components/icons/base/ChapterNext'
import ChapterPrevious from 'grommet/components/icons/base/ChapterPrevious'
import Next from 'grommet/components/icons/base/Next'
import Previous from 'grommet/components/icons/base/Previous'
import Order from './Order'

function navLinks(links, page, onPageLinkClick) {
    var navLinks = [];
    if ("first" in links && page.number > 0) {
        navLinks.push(<Anchor key="first" href="" icon={<ChapterPrevious />} onClick={(e)=>onPageLinkClick('first', e)} />);
    }
    if ("prev" in links) {
        navLinks.push(<Anchor key="prev" href="" icon={<Previous />} onClick={(e)=>onPageLinkClick('prev', e)} />);
    }
    if ("next" in links) {
        navLinks.push(<Anchor key="next" href="" icon={<Next />} onClick={(e)=>onPageLinkClick('next', e)} />);
    }
    if ("last" in links && page.number < (page.totalPages-1)) {
        navLinks.push(<Anchor key="last" href="" icon={<ChapterNext />} onClick={(e)=>onPageLinkClick('last', e)} />);
    }
    return navLinks;
}

const OrderList = ({ items, _links, page, onPageLinkClick }) => (
        <div>
        <table>
            <thead>
                <tr>
                <th>Name</th>
                <th>Status</th>
                <th>Ordered Date</th>
                <th></th>
                </tr>
            </thead>
            <tbody>
                { items.map(order =>
                    <Order key={order._links.self.href} {...order} />
                ) }
            </tbody>
        </table>
        <div>
            {navLinks(_links, page, onPageLinkClick)}
        </div>
    </div>
)

// TODO: is it necessary to declare the shape of an order here?  or could it
// be defined in Order.js and re-used?
OrderList.propTypes = {
    _links: PropTypes.object.isRequired, // TODO: more explicit shape
    page: PropTypes.shape({
        size : PropTypes.number.isRequired,
        totalElements : PropTypes.number.isRequired,
        totalPages : PropTypes.number.isRequired,
        number : PropTypes.number.isRequired,
    }),
    items: PropTypes.arrayOf(
               PropTypes.shape(Order.propTypes)
    ).isRequired,
    onPageLinkClick: PropTypes.func.isRequired,
}


export default OrderList

