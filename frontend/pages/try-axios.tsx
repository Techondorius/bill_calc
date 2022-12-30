import axios, { AxiosResponse } from 'axios'
import { useState } from 'react'

const baseURL = 'http://localhost:8080'

const request = () => {
  interface indexRes {
    message?: string
  }
  const [res, setRes] = useState('')
  const fetchData = async () => {
    const msg = await axios.get<indexRes>(baseURL)
    setRes(msg.data.message)
  }
  fetchData()

  return (
    <>
      <h1>{res}</h1>
      {/* <button>fetchData</button> */}
    </>
  )
}

export default request
