import styles from '../styles/Home.module.css';
import { Grommet, Button, Box } from 'grommet';

import styled from 'styled-components';

const StyledCenter = styled.div`
  position: fixed;
  top: 50%;
  left: 50%;
  margin: auto;
  font-size: 50px;
  transform: translate(-50%, -50%);
  color: #77A2BC;
  padding: 40px;
  border: 2px dashed #77A2BC;
  &:hover {
    border: 2px dashed white;
    text-shadow: 0 0 30px #728CA7;
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
