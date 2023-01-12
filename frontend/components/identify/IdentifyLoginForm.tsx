import { TextField } from '@mui/material'
import axios, { AxiosError } from 'axios'
import { MyButton } from 'components/button/MyButton'
import { useState } from 'react'
import { useRouter } from 'next/router'

const IdentifyLoginForm = () => {
  const [uid, setUid] = useState('')
  const [pwd, setPwd] = useState('')
  const [errMsg, setErrMsg] = useState('')
  interface loginRes {
    token: string
  }
  interface loginResError {
    error: string
  }
  const router = useRouter()
  const submit = async () => {
    await console.log(uid)
    await console.log(pwd)
    axios
      .post<loginRes>(
        'http://localhost:8080/login',
        {
          userid: uid,
          password: pwd,
        },
        {
          withCredentials: true,
        },
      )
      .then(() => {
        setErrMsg('')
        router.push('/bills')
      })
      .catch((e: AxiosError<loginResError>) => {
        console.log(e.response.data.error)
        console.log(e)
        setErrMsg(e.response.data.error)
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
        <p>{errMsg}</p>
      </div>
    </>
  )
}

export default IdentifyLoginForm
