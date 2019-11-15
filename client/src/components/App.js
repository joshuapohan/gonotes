import React from 'react';

class App extends React.Component {
    constructor(props){
        super(props);
        this.ws = null;
    }

    fetchRoom = (e) => {
        e.preventDefault();
        let room = document.querySelector("#room").value;
        this.ws = new WebSocket("ws://127.0.0.1:8080/ws?room=" + room);
    }

    render(){
        return(
            <div>
                <input id="room" type="text"/>
                <button onClick={this.fetchRoom}>Enter</button>
            </div>
        );
    }
}

export default App;