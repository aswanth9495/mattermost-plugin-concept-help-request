export const OPEN_CHR_MODAL = 'chr_plugin/chr_modal/OPEN_MODAL';
export const CLOSE_CHR_MODAL = 'chr_plugin/chr_modal/CLOSE_MODAL';

export function openChrModal({postId}) {
    return (dispatch) => {
        dispatch({type: OPEN_CHR_MODAL, payload: postId});
    };
}

export function closeChrModal() {
    return (dispatch) => {
        dispatch({type: CLOSE_CHR_MODAL});
    };
}

export default function reducer(
    state = {
        postId: false,
    },
    action,
) {
    switch (action.type) {
    case OPEN_CHR_MODAL: {
        return {
            ...state,
            postId: action.payload,
        };
    }
    case CLOSE_CHR_MODAL: {
        return {
            ...state,
            postId: false,
        };
    }
    default:
        return state;
    }
}