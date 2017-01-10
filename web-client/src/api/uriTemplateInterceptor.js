https://github.com/spring-guides/tut-react-and-spring-data-rest/blob/89c5b5191c5e36785b670abe195d35b5e405b223/hypermedia/src/main/resources/static/api/uriTemplateInterceptor.js
define(function(require) {
    'use strict';

    var interceptor = require('../../node_modules/rest/interceptor');

    return interceptor({
        request: function (request /*, config, meta */) {
            /* If the URI is a URI Template per RFC 6570 (http://tools.ietf.org/html/rfc6570), trim out the template part */
            if (request.path.indexOf('{') === -1) {
                return request;
            } else {
                request.path = request.path.split('{')[0];
                return request;
            }
        }
    });
});
