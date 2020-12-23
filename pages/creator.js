import { Grommet, Box, Grid, Heading } from 'grommet';
import React, { useRef, useEffect } from 'react';

import styled from 'styled-components';

const StyledCanvas = styled.canvas`
  height: 100%;
  width: 100%;
`;

const drawImage = (context, imgSrc) => {
  const img = new Image();

  img.src = imgSrc;
  img.onload = () => {
    context.drawImage(
      img,
      0.5,
      0.5,
      context.canvas.width * 0.9,
      img.height * (context.canvas.width / img.width) * 0.9
    );
  };
};

const Canvas = (props) => {
  const canvasRef = useRef(null);

  useEffect(() => {
    const con = document.getElementById('canvasBox');
    const canvas = canvasRef.current;
    console.log(canvas.width, canvas.height, con.clientWidth, con.clientHeight);
    canvas.width = con.clientWidth;
    canvas.height = con.clientHeight;
    const context = canvas.getContext('2d');
    context.imageSmoothingEnabled = true;
    //Our first draw
    context.fillStyle = '#000000';
    console.log(context.canvas.width, context.canvas.height);
    drawImage(context, 'assets/result.png');
  }, []);

  return <StyledCanvas ref={canvasRef} {...props} />;
};

const Creator = (props) => {
  return (
    <Grommet full>
      <Grid
        rows={['175px', '1fr']}
        columns={['1fr', '80%', '1fr']}
        areas={[
          ['lside', 'header', 'rside'],
          ['lside', 'main', 'rside'],
          ['lside', 'main', 'rside'],
        ]}
        align="stretch"
      >
        <Box gridArea="header" background="#84c1c4">
          <Heading
            style={{ color: '#ffffff', margin: 'auto', fontWeight: 400 }}
          >
            avi-maker
          </Heading>
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
