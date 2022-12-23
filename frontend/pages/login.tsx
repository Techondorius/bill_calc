import IdentifyLoginForm from 'components/identify/IdentifyLoginForm'

function login() {
  return (
    <>
      <div className='flex justify-center w-full'>
        <div className='flex-col my-3 w-5/6 px-3 py-3 sm:w-1/2'>
          <IdentifyLoginForm />
        </div>
      </div>
    </>
  )
}

export default login
