// https://github.com/spring-guides/tut-react-and-spring-data-rest/blob/89c5b5191c5e36785b670abe195d35b5e405b223/hypermedia/src/main/resources/static/api/uriListConverter.js
define(function() {
    'use strict';

    /* Convert a single or array of resources into "URI1\nURI2\nURI3..." */
    return {
        read: function(str /*, opts */) {
            return str.split('\n');
        },
        write: function(obj /*, opts */) {
            // If this is an Array, extract the self URI and then join using a newline
            if (obj instanceof Array) {
                return obj.map(function(resource) {
                    return resource._links.self.href;
                }).join('\n');
            } else { // otherwise, just return the self URI
                return obj._links.self.href;
            }
        }
    };
});
