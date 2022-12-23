import { Button, TextField } from '@mui/material'
import { MyButton } from 'components/button/MyButton'
import { useState } from 'react'

const IdentifyLoginForm = () => {
  const [uid, setUid] = useState('')
  const [pwd, setPwd] = useState('')
  const submit = () => {
    console.log(uid)
    console.log(pwd)
  }
  return (
    <>
      <div className='space-y-3'>
        <h1 className='text-3xl'>Login</h1>
        <TextField
          fullWidth
          required
          label='UserID'
          value={uid}
          onChange={(newUid) => {
            setUid(newUid.target.value)
          }}
        />
        <TextField
          fullWidth
          required
          id='outlined-required'
          label='Password'
          value={pwd}
          type='password'
          onChange={(newPwd) => {
            setPwd(newPwd.target.value)
          }}
        />
        <MyButton type='main' className='rounded-full w-full'>LOGIN</MyButton>
        {/* <MyButton type='main'>asdf</MyButton> */}
      </div>
    </>
  )
}

export default IdentifyLoginForm
