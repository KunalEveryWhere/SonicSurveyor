import React, { useState, useRef } from 'react';
import RecordSoundStartIcon from "../assets/recordSoundStartIcon.png"
import RecordSoundStopIcon from "../assets/recordSoundStopIcon.png"

const RecordSoundButton = ({onSendAverageNoiseDecibel}) => {
  const audioRef = useRef(null);
  const [isRecording, setIsRecording] = useState(false);
  const [averageNoiseDecibel, setAverageNoiseDecibel] = useState(0);
  const [mediaRecorder, setMediaRecorder] = useState(null);
  const chunksRef = useRef([]);

  const handleRecord = async () => {
    try {
      if (!isRecording) {
        const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
        const mediaRecorderInstance = new MediaRecorder(stream);
        setMediaRecorder(mediaRecorderInstance);
        chunksRef.current = [];

        mediaRecorderInstance.ondataavailable = (e) => {
          if (e.data.size > 0) {
            chunksRef.current.push(e.data);
          }
        };

        mediaRecorderInstance.onstop = () => {
          const audioBlob = new Blob(chunksRef.current, { type: 'audio/wav' });
          const audioURL = URL.createObjectURL(audioBlob);
          audioRef.current.src = audioURL;

          const audioContext = new AudioContext();
          const reader = new FileReader();

          reader.onloadend = () => {
            audioContext.decodeAudioData(reader.result, (buffer) => {
              const channelData = buffer.getChannelData(0); // Assuming mono audio
              const squaredValues = channelData.map((sample) => sample * sample);
              const meanSquare = squaredValues.reduce((sum, sample) => sum + sample, 0) / squaredValues.length;
              const rootMeanSquare = Math.sqrt(meanSquare);
              const reference = 0.000015; // Reference amplitude (requires the average decibel value is consistently higher than expected, you could try increasing the reference amplitude to see if it brings the average decibel value closer to the expected value.)
              const decibel = 20 * Math.log10(rootMeanSquare / reference);
              const offset = -0; // Offset value in dB
              setAverageNoiseDecibel((decibel + offset).toFixed(2));
              onSendAverageNoiseDecibel((decibel + offset).toFixed(2));
            });
          };

          reader.readAsArrayBuffer(audioBlob);
        };

        mediaRecorderInstance.start();
        setIsRecording(true);
      } else {
        mediaRecorder.stop();
        setIsRecording(false);
      }
    } catch (error) {
      console.error('Error accessing microphone:', error);
      console.log(averageNoiseDecibel);
    }
  };

  return (
    <div>
      <button onClick={handleRecord} className='recordSoundButton' style={{ backgroundColor: isRecording ? '#FF5FAA' : 'white' }}>{isRecording ? 
      <img src={RecordSoundStopIcon} alt="Record Sound Stop Icon" width="48" height="auto"/> 
      : 
      <img src={RecordSoundStartIcon} alt="Record Sound Start Icon" width="48" height="auto"/> 
      }</button>
      <audio ref={audioRef} controls style={{display: 'none'}}/>
    </div>
  );
};

export default RecordSoundButton;
