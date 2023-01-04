import axios from 'axios'
import { MyButton } from 'components/button/MyButton'
import { useState } from 'react'

const baseURL = 'http://localhost:8080/rand'

const Request = () => {
  interface indexRes {
    message?: string
  }
  const [res, setRes] = useState('')
  const fetchData = async () => {
    const msg = await axios.get<indexRes>(baseURL)
    setRes(msg.data.message)
  }
  const a = () => fetchData()

  return (
    <>
      <h1>{res}</h1>
      <MyButton type='main' onClick={a}>
        fetchData
      </MyButton>
    </>
  )
}

export default Request
