import React from 'react';
import logo from './logo.svg';
import { Card, Icon, Image } from 'semantic-ui-react'
import './App.css';

function App() {
  return (
    <div className="App">
      <Card>
        <Image src='https://react.semantic-ui.com/images/avatar/large/matthew.png' wrapped ui={false} />
        <Card.Content>
          <Card.Header>Matthew</Card.Header>
          <Card.Meta>
            <span className='date'>Joined in 2015</span>
          </Card.Meta>
          <Card.Description>
            Matthew is a musician living in Nashville.
      </Card.Description>
        </Card.Content>
        <Card.Content extra>
          <a>
            <Icon name='user' />
            22 Friends
      </a>
        </Card.Content>
      </Card>
    </div>
  );
}

export default App;
