import styles from '../styles/Home.module.css';
import { Grommet, Button, Box } from 'grommet';

import styled from 'styled-components';

const StyledCenter = styled.div`
  position: fixed;
  top: 40%;
  left: 50%;
  margin: auto;
  font-size: 50px;
  transform: translate(-50%, -50%);
  color: #646b9a;
  border-radius: 10px;
  padding: 50px;
  border: 2px dashed #646b9a;
  &:hover {
    border: 2px dashed white;
    text-shadow: 0 0 20px white;
    color: white;
  }
`;

export default function Home() {
  return (
    <Grommet>
      <StyledCenter>
        <a href="/creator">welcome</a>
      </StyledCenter>
    </Grommet>
  );
}
