import { Grommet, Box, Grid, Heading } from 'grommet';
import React, { useRef, useEffect } from 'react';

import styled from 'styled-components';

const StyledCanvas = styled.canvas`
  height: 100%;
  width: 100%;
`;

const StyledFooter = styled(Heading)`
  color: #ffffff;
  font-weight: 400;
  text-align: center;
  margin: auto;
  font-size: 15px;
`
 
const drawImage = (context, imgSrc) => {
  const img = new Image();
  img.onload = () => {
    const imgWidth = img.naturalWidth;
    const imgHeight = img.naturalHeight;
    const canvasWidth = context.canvas.width;
    const canvasHeight = context.canvas.height;
    // scale based on the limiting factor
    let newWidth = imgWidth;
    let newHeight = imgHeight;
    if (canvasWidth < imgWidth && canvasHeight < imgHeight) {
      if (canvasWidth < canvasHeight) {
        newWidth = canvasWidth;
        newHeight = imgHeight * (canvasWidth / imgWidth);
      } else {
        newWidth = imgWidth * (canvasHeight / imgHeight);
        newHeight = canvasHeight;
      }
    } else if (canvasWidth < imgWidth) {
      newWidth = canvasWidth;
      newHeight = imgHeight * (canvasWidth / imgWidth);
    } else if (canvasHeight < imgHeight) {
      newWidth = imgWidth * (canvasHeight / imgHeight);;
      newHeight = canvasHeight;;
    }
    console.log(
      'image width and height:',
      imgWidth,
      imgHeight,
      'canvas width and height:',
      canvasWidth,
      canvasHeight,
      'new width and height:',
      newWidth,
      newHeight
    );
    context.drawImage(
      img,
      (canvasWidth - newWidth) / 2,
      (canvasHeight - newHeight) / 2,
      newWidth,
      newHeight
    );
  };
  img.src = imgSrc;
};

const Canvas = (props) => {
  const canvasRef = useRef(null);

  useEffect(() => {
    const parent = document.getElementById('canvasBox');
    const canvas = canvasRef.current;
    console.log(canvas.width, canvas.height, parent.clientWidth, parent.clientHeight);
    canvas.width = parent.clientWidth;
    canvas.height = parent.clientHeight;
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
        rows={['150px', '1fr', '20px']}
        columns={['1fr', '80%', '1fr']}
        areas={[
          ['lside', 'header', 'rside'],
          ['lside', 'main', 'rside'],
          ['lside', 'main', 'rside'],
          ['footer', 'footer', 'footer'],
        ]}
        align="stretch"
        gap="small"
      >
        <Box gridArea="header">
          <Heading
            style={{
              color: '#ffffff',
              margin: 'auto',
              fontWeight: 400,
              border: '2px dashed white',
              padding: '15px',
              borderRadius: '10px',
              fontSize: '30px',
            }}
          >
                avatar creator
          </Heading>
        </Box>
        <Box
          gridArea="main"
          style={{
            color: '#ffffff',
            fontWeight: 400,
            border: '3px dashed #84c1c4',
            padding: '15px',
            borderRadius: '10px',
          }}
        >
          <Grid
            columns={['1fr', '1fr']}
            rows={['100%']}
            areas={[['avatar', 'customization']]}
            align="stretch"
            style={{ height: '75vh' }}
          >
            <Box
              id="canvasBox"
              gridArea="avatar"
              background="white"
              style={{
                color: '#ffffff',
                fontWeight: 400,
                borderRadius: '10px',
              }}
            >
              <Canvas />
            </Box>
          </Grid>
        </Box>
        <Box
          gridArea="footer"
          style={{
            fontWeight: 400,
          }}
        >
          <StyledFooter>made with ♥ in san francisco</StyledFooter>
        </Box>
      </Grid>
    </Grommet>
  );
};

export default Creator;
