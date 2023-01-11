import { Button, TextField } from '@mui/material'
import axios, { AxiosError } from 'axios'
import { MyButton } from 'components/button/MyButton'
import { useState } from 'react'

const IdentifyLoginForm = () => {
  const [uid, setUid] = useState('')
  const [pwd, setPwd] = useState('')
  interface loginRes {
    token: string
  }
  interface loginResError {
    error: string
  }
  const submit = async () => {
    await console.log(uid)
    await console.log(pwd)
    axios
      .post<loginRes>('http://localhost:8080/login', { userid: uid, password: pwd })
      .then((res) => {
        console.log(res.data.token)
      })
      .catch((e: AxiosError<loginResError>) => {
        console.log(e.response)
      })
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
        <MyButton type='main' className='rounded-full w-full' onClick={submit}>
          LOGIN
        </MyButton>
      </div>
    </>
  )
}

export default IdentifyLoginForm
