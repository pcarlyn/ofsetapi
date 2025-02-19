import './App.css'
import { Button, Highlight } from "@chakra-ui/react"
import {
  NativeSelectField,
  NativeSelectRoot,
} from "./components/ui/native-select"
import { Toaster, toaster } from "./components/ui/toaster"
import { useState, useRef, useCallback, useEffect } from 'react'
import { Field, Input } from "@chakra-ui/react"



function App() {
  const [inputValue, setInputValue] = useState('');
  const [num, setNum] = useState(0)
  const [selector, setSelector] = useState(1)
  const [log, setLog] = useState("")
  const [placeholder, setPlaceholder] = useState(0);
  const [host, setHost] = useState('');
  const [error, setError] = useState('');


  useEffect(() => {
    console.log('placeholder изменился:', placeholder);
  }, [placeholder]);

  useEffect(() => {
    console.log('error изменился:', error);
  }, [error]);

  const calculateHost = (value: string | number, variable: number): string =>{
    let numericValue: number;
  
    if (value === 'of') {
      numericValue = 10;
    } else if (value === 'lt') {
      numericValue = 2;
    } else if (typeof value === 'number') {
      numericValue = value;
    } else {
      throw new Error('Invalid value input');
    }
  
    if (variable > 255) {
      const multiple = Math.floor(variable / 255);
      numericValue += multiple;
      variable = variable % 256;
    }

    if (numericValue > 255) {
      throw new Error('Invalid value input');
    }
  
    return `10.10.${numericValue}.${variable}`;
  }
  
  const handleCalculateHost = (value: string | number, variable: number) => {
    try {
      const calculatedHost = calculateHost(value, variable);
      setHost(calculatedHost);
      setError('');
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message);
      } else {
        setError('Произошла неизвестная ошибка'); 
      }
      setHost('');
    }
  };
  
  
  const timeoutRef = useRef<number | null>(null);

  const debouncedFetchData = useCallback((newPlaceholder: number) => {
    if (timeoutRef.current) {
      clearTimeout(timeoutRef.current);
    }

    timeoutRef.current = window.setTimeout(() => {
      fetchData(newPlaceholder);
      timeoutRef.current = null;
    }, 500);
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value;
    setInputValue(newValue);
    const newPlaceholder = Number(newValue);
    setPlaceholder(newPlaceholder);
    handleCalculateHost("of", newPlaceholder)
    debouncedFetchData(newPlaceholder);
  };

  const handleInputKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      const newPlaceholder = Number(inputValue); 
      setPlaceholder(newPlaceholder);
      handleCalculateHost("of", newPlaceholder)
      fetchData(newPlaceholder);
    }
  };

  const fetchData = async (newPlaceholder: number) => {

    const newHost = calculateHost("of", newPlaceholder);
    try {
      const response = await fetch("http://" + newHost + ":1323/control/get-number");

      if (!response.ok) {
        setNum(0);
        return;
      }

      const result = await response.json();
      const parsedNum = Number(result);

      if (isNaN(parsedNum)) {
        console.error("Получено нечисловое значение:", result);
        setNum(0);
        return;
      }

      setNum(parsedNum);
    } catch (error) {
      console.error("Ошибка при запросе:", error);
      setNum(0);
    }
  };
  
  
  return (
    <>
    <div className="flex justify-center items-center">
      <div className="w-72">
        <Field.Root invalid>
          <Field.Label>of***.offset-partners.ru</Field.Label>
          <Input 
            placeholder={"444"}
            value={inputValue}
            onChange={handleInputChange}
            onKeyDown={handleInputKeyDown}
          />
          
          <Field.ErrorText>Максимум трехзначные числа</Field.ErrorText>
        </Field.Root>
      </div>
    </div>

    <br></br>

    <div className="flex">
      <div className="w-100 flex-12  ...">
        <Highlight query="Offset" styles={{ fontWeight: "semibold" }}>
        {`Offset-${num} Host address-${host}`}
        </Highlight>
      </div>
    </div>
   
    <br></br>
    <div className="flex justify-center items-center">
      <div className="w-72">
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
        var baseUrl = 'http://' + host + ':1323'
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
          </div>
          </div>
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
