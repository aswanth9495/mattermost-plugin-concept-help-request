import React from 'react';
import {useSelector, shallowEqual, useDispatch} from 'react-redux';

import {pluginId} from '../utils/constants';

import {closeChrModal} from '../stores/chrModal';

import ChrModal from './ChrModal';

import '../index.scss';

function Root() {
    const {
        postId,
    } = useSelector((state) => {
        const currentState = state ? state['plugins-' + pluginId] : {};
        return {
            postId: currentState.chrModalStore.postId,
        };
    }, shallowEqual);

    const dispatch = useDispatch();

    return (
        <ChrModal
            show={Boolean(postId)}
            onHide={() => {
                dispatch(closeChrModal());
            }}
            style={{opacity: 1}}
        />
    );
}

export default Root;