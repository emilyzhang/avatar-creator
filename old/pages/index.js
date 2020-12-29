import styles from '../styles/Home.module.css';
// import styled from 'styled-components';
import dynamic from 'next/dynamic';

// const Stage = dynamic(
//   () => {
//     return import('react-pixi-fiber').then((pixi) => pixi.Stage);
//   },
//   { ssr: false }
// );
// const Sprite = dynamic(
//   () => {
//     return import('react-pixi-fiber').then((pixi) => pixi.Sprite);
//   },
//   { ssr: false }
// );
// import styled from 'styled-components';
// import { Sprite, Stage } from 'react-pixi-fiber';

const StyledStage = styled(Stage)`
  position: fixed;
  top: 50%;
  left: 25%;
  margin: auto;
  font-size: 50px;
  transform: translate(-50%, -50%);
  color: #77a2bc;
  padding: 40px;
  border: 2px dashed #77a2bc;
  &:hover {
    border: 2px dashed white;
    text-shadow: 0 0 30px #728ca7;
    color: white;
  }
`;

function Bunny() {
  return <Sprite />;
}

const Canvas = (props) => {
  const canvasRef = useRef(null);

  useEffect(() => {}, []);

  return <canvas ref={canvasRef} {...props} />;
};


export default function Home() {
  return (
    <StyledStage
      options={{ backgroundColor: 'white', height: 500, width: 500 }}
    ></StyledStage>
  );
}
