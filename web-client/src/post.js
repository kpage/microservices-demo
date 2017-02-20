// TODO: de-dupe logic with follow.js
module.exports = function post(api, rootPath, relArray) {
    // TODO: allow posting to the root too?
    var root = api({
        method: 'GET',
        path: rootPath
    });

    var [lastItem] = relArray.slice(-1)
    return relArray.reduce(function(root, arrayItem) {
        var rel = typeof arrayItem === 'string' ? arrayItem : arrayItem.rel;
        var isPost = (arrayItem === lastItem);
        return traverseNext(root, rel, arrayItem, isPost);
    }, root);

    function traverseNext(root, rel, arrayItem, isPost) {
        return root.then(function (response) {
            if (hasEmbeddedRel(response.entity, rel)) {
                return response.entity._embedded[rel];
            }

            if(!response.entity._links) {
                return [];
            }

            if (typeof arrayItem === 'string') {
                return api({
                    method: isPost ? 'POST' : 'GET',
                    path: response.entity._links[rel].href,
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    entity: arrayItem.entity
                });
            } else {
                return api({
                    method: isPost ? 'POST' : 'GET',
                    path: response.entity._links[rel].href,
                    params: arrayItem.params,
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    entity: arrayItem.entity
                });
            }
        });
    }

    function hasEmbeddedRel(entity, rel) {
        return entity._embedded && entity._embedded.hasOwnProperty(rel);
    }
};
