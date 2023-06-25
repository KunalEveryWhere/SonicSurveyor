import LoadingIcon from "../assets/loadingIcon.gif";

const Loading = () => {
  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '80vh',
        flexDirection: 'column',
        gap: '0px'
      }}
    >
    <img src={LoadingIcon} alt="Loading Icon"/>
      <h2>Loading...</h2>
      <h4>Please wait a moment, this might take some time ⏱️</h4>
      <p>🧮 Heavy Calculation Happening...</p>
    </div>
  );
};

export default Loading;