import manifest from './manifest';
import {openChrModal} from './stores/chrModal';
import reducer from './stores/index';
import Root from './components/Root';

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
        // eslint-disable-next-line no-console
        console.log(store);

        // @see https://developers.mattermost.com/extend/plugins/webapp/reference/
        registry.registerRootComponent(Root);
        registry.registerPostDropdownMenuAction(
            'Raise CHR',
            (postId) => (store.dispatch(openChrModal({postId}))),
        );
        registry.registerReducer(reducer);
    }
}

window.registerPlugin(manifest.id, new Plugin());
