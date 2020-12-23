import { Grommet, Box, Grid } from 'grommet';
import React, { useRef, useEffect } from 'react';

import styled from 'styled-components';

const StyledCanvas = styled.canvas`
  height: 100%;
  width: 100%;
`;

const Canvas = (props) => {
  const canvasRef = useRef(null);

  useEffect(() => {
    const con = document.getElementById('canvasBox');
    const canvas = canvasRef.current;
    console.log(canvas.width, canvas.height, con.clientWidth, con.clientHeight);
    canvas.width = con.clientWidth;
    canvas.height = con.clientHeight;
    const context = canvas.getContext('2d');
    context.imageSmoothingEnabled = false;
    //Our first draw
    context.fillStyle = '#000000';
    console.log(context.canvas.width, context.canvas.height);
    context.fillRect(0.5, 0.5, 20, 5000);
    var img = new Image();
    img.src = 'assets/result.png';
    img.onload = () => {
      context.drawImage(img, 0, 0, context.canvas.width, context.canvas.height);
    };
  }, []);

  return <StyledCanvas ref={canvasRef} {...props} />;
};

const Creator = (props) => {
  return (
    <Grommet full>
      <Grid
        rows={['200px', '1fr']}
        columns={['1fr', '80%', '1fr']}
        areas={[
          ['lside', 'header', 'rside'],
          ['lside', 'main', 'rside'],
          ['lside', 'main', 'rside'],
        ]}
        align="stretch"
      >
        <Box gridArea="header" background="white">
          "hello world"
        </Box>
        <Box gridArea="main" background="white">
          <Grid
            columns={['1fr', '1fr']}
            rows={['100%']}
            areas={[['avatar', 'customization']]}
            align="stretch"
            style={{ height: '75vh' }}
          >
            <Box id="canvasBox" gridArea="avatar" background="white">
              <Canvas />
            </Box>
          </Grid>
        </Box>
        <Box gridArea="footer" background="white">
          b
        </Box>
      </Grid>
    </Grommet>
  );
};

export default Creator;
