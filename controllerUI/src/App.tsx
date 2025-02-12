import './App.css'
import { Button, Highlight } from "@chakra-ui/react"
import {
  NativeSelectField,
  NativeSelectRoot,
} from "./components/ui/native-select"
import { Toaster, toaster } from "./components/ui/toaster"
import { useState, useEffect } from 'react'

function App() {
  const [num, setNum] = useState(0)
  const [selector, setSelector] = useState(1)

  useEffect(() => {
    async function fetchData() {
      const response = await fetch("http://10.10.10.159:1323/control/get-number");
      const result = await response.json();
      setNum(Number(result));
    }

    fetchData();}
  )
  
  return (
    <>
    <div className="flex">
      <div className="w-100 flex-12  ...">
        <Highlight query="Offset" styles={{ fontWeight: "semibold" }}>
          Offset-{}
        </Highlight>
        {String(num)}
      </div>
    </div>
   
    <br></br>
      <NativeSelectRoot>
        <NativeSelectField onChange={(e) => setSelector(Number(e.target.value))}>
          <option value="1">Очиститть очередь</option>
          <option value="2">Выполнить тестовую печать</option>
        </NativeSelectField>
      </NativeSelectRoot>
      <br></br>
      <Button
      variant="outline"
      size="sm"
      onClick={() => {
        const promise = new Promise<void>((resolve) => {
          setTimeout(() => resolve(), 5000)
        })
        // setResponse("Ooops")
        var url = 'http://10.10.10.159:1323/control/order-clean'
        if (selector === 2) {url = 'http://10.10.10.159:1323/control/print-test'}
        fetch(url, { 
          method: 'GET', 
          headers: new Headers({
              'Authorization': 'fdasdfsd'
          }),
          mode: 'cors'
        }).then(async (res) => {
          console.log(await res.json().toString)
        })
        toaster.promise(promise, {
          success: {
            title: "Success",
            description: "response",
          },
          error: {
            title: "Failed",
            description: "response",
          },
          loading: { title: "Uploading...", description: "Please wait" },
        })
        
      }}>Выполнить</Button>
      <Toaster />
    </>
  )
}

export default App
