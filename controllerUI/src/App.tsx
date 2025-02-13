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
  const [log, setLog] = useState("")

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
          <option value="1">Очистить очередь</option>
          <option value="2">Выполнить тестовую печать</option>
          <option value="3">Логи lpf info</option>
          <option value="4">Логи lpf error</option>
          <option value="5">speedtest</option>
          <option value="6">speedtest --secure</option>
        </NativeSelectField>
      </NativeSelectRoot>
      <br></br>
      <Button className="mb-4, pb-10"
      variant="outline"
      size="sm"
      onClick={() => {
        // const promise = new Promise<void>((resolve) => {
        //   setTimeout(() => resolve(), 5000)
        // })
        var baseUrl = 'http://localhost:1323'
        var url = baseUrl + '/control/order-clean'
        type ResponseData = { status: number; text: string };
        if (selector === 2) {url = baseUrl + '/control/print-test'}
        if (selector === 3) {url = baseUrl + '/control/get-lpf-info'}
        if (selector === 4) {url = baseUrl + '/control/get-lpf-error'}
        if (selector === 5) {url = baseUrl + '/control/run-speedtest'}
        if (selector === 6) {url = baseUrl + '/control/run-speedtest-secure'}
        const prom = fetch(url, { 
          method: 'GET', 
          headers: new Headers({
              'Authorization': '8a1829a0b9094e6392242dcb242a5b1e'
          }),
          mode: 'cors'
        }).then(async (res) => {
          const text = await res.text();
          setLog(text);
          if (!res.ok) {
            throw new Error(`Ошибка ${res.status}: ${res.statusText}\n${text}`);
          }
          return { status: res.status, text };
        })
        toaster.promise(prom, {
          success: (data: ResponseData) => ({
            title: 'Success',
            description: `Status: ${data.status}`,
          }),
          error: (err) => ({
            title: 'Failed',
            description: err.message,
          }),
          loading: {
            title: 'Loading...',
            description: 'Запрос выполняется...',
          },
        })
        
      }}>Выполнить</Button>
      <br></br>
      <br></br>
      <Toaster />
      <div className="bg-black text-green-400 font-mono p-4 rounded-lg shadow-lg max-h-96 overflow-y-auto border border-gray-600">
        <pre style={{ textAlign: 'left' }} className="whitespace-pre-wrap">{log}</pre>
      </div>

    </>
  )
}

export default App
