'use strict';

// https://github.com/spring-guides/tut-react-and-spring-data-rest/blob/89c5b5191c5e36785b670abe195d35b5e405b223/hypermedia/src/main/resources/static/client.js
var rest = require('../node_modules/rest');
var defaultRequest = require('../node_modules/rest/interceptor/defaultRequest');
var mime = require('../node_modules/rest/interceptor/mime');
var errorCode = require('../node_modules/rest/interceptor/errorCode');
var baseRegistry = require('../node_modules/rest/mime/registry');
var uriTemplateInterceptor = require('./api/uriTemplateInterceptor');

var registry = baseRegistry.child();

registry.register('text/uri-list', require('./api/uriListConverter'));
registry.register('application/hal+json', require('../node_modules/rest/mime/type/application/hal'));

module.exports = rest
        .wrap(mime, { registry: registry })
        .wrap(uriTemplateInterceptor)
        .wrap(errorCode)
        .wrap(defaultRequest, { headers: { 'Accept': 'application/hal+json' }});
