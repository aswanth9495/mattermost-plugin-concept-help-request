import React from 'react';
const {
    Modal,
    Button,
    FormGroup,
    ControlLabel,
    FormControl,
    Form,
    Radio,
} = window.ReactBootstrap;

function topicUi() {
    return (
        <FormGroup controlId='chr-topic'>
            <ControlLabel>Topic</ControlLabel>
            <FormControl
                componentClass='select'
                placeholder='Select topic'
            >
                <option value='select'>Time Complexity</option>
                <option value='other'>...</option>
            </FormControl>
        </FormGroup>
    );
}

function subtopicUi() {
    return (
        <FormGroup
            controlId='chr-subtopic'
            className='chr-modal__subtopic'
        >
            <FormControl
                componentClass='select'
                placeholder='Select a Concept'
            >
                <option value='select'>Select a Concept</option>
                <option value='other'>...</option>
            </FormControl>
        </FormGroup>
    );
}

function doubtUi() {
    return (
        <FormGroup
            controlId='chr-doubt'
            className='chr-modal__doubt'
        >
            <FormControl
                componentClass='textarea'
                placeholder='Write your doubt here'
            />
        </FormGroup>
    );
}

function selectedtimeSlotUi() {
    return (
        <div className='chr-modal__time-slot'>
            Here the time-slots go
        </div>
    );
}

function dateUi() {
    return (
        <>
            <ControlLabel>Choose Date</ControlLabel>
            <FormGroup>
                <Radio
                    name='radioGroup'
                    inline={true}
                >
                    Fri 4, Sep
                </Radio>{' '}
                <Radio
                    name='radioGroup'
                    inline={true}
                >
                    Sat, 5 Sep
                </Radio>
            </FormGroup>
        </>
    );
}

function timeUi() {
    return (
        <FormGroup controlId='chr-topic'>
            <ControlLabel>Topic</ControlLabel>
            <FormControl
                componentClass='select'
                placeholder='Select topic'
            >
                <option value=''>Select time Slot</option>
                <option value='other'>9:00 AM</option>
                <option value='other'>10:00 AM</option>
                <option value='other'>11:00 AM</option>
                <option value='other'>12:00 AM</option>
            </FormControl>
        </FormGroup>
    );
}

function ChrModal(props) {
    return (
        <Modal
            {...props}
            size='lg'
            aria-labelledby='contained-modal-title-vcenter'
            centered={true}
            className='chr-modal'
        >
            <Modal.Header closeButton={true}>
                <Modal.Title id='contained-modal-title-vcenter'>
                    Create Concept Help Request
                </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <h3>Having trouble understanding a Concept/Theory?</h3>
                <h5>Please provide us with some info on your doubt and Get on a 1-1 live call with TA/Peer</h5>
                <Form className='chr-modal__form'>
                    {topicUi()}
                    <div className='chr-modal__doubt-section'>
                        <ControlLabel>What are your doubts ?</ControlLabel>
                        <div className='chr-modal__doubt-wrap'>
                            {subtopicUi()}
                            {doubtUi()}
                        </div>
                    </div>
                    {dateUi()}
                    {timeUi()}
                    {selectedtimeSlotUi()}
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button bsStyle='primary'>Submit</Button>
                <Button onClick={props.onHide}>Close</Button>
            </Modal.Footer>
        </Modal>
    );
}

export default ChrModal;

