import './App.css'
import { Button, Highlight } from "@chakra-ui/react"
import {
  NativeSelectField,
  NativeSelectRoot,
} from "./components/ui/native-select"
import { Toaster, toaster } from "./components/ui/toaster"
import { useState, useRef, useCallback, useEffect } from 'react'
import { Field, Input } from "@chakra-ui/react"
import image from './assets/LISTOK.png';



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
    handleCalculateHost("lt", newPlaceholder)
    debouncedFetchData(newPlaceholder);
  };

  const handleInputKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      const newPlaceholder = Number(inputValue); 
      setPlaceholder(newPlaceholder);
      handleCalculateHost("lt", newPlaceholder)
      fetchData(newPlaceholder);
    }
  };

  const fetchData = async (newPlaceholder: number) => {

    const newHost = calculateHost("lt", newPlaceholder);
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
    <br></br>
    <div className="flex justify-center items-center">
    <img src={image} alt="Listok" className="w-48 h-auto" />
    </div>
    
    <div className="flex justify-center items-center">
      <div className="w-72">
        <Field.Root invalid>
          <Field.Label>lt***.listcopy.local</Field.Label>
          <Input 
            placeholder={"111"}
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
        {`LISTOK-${num} Host address-${host}`}
        </Highlight>
      </div>
    </div>
   
    <br></br>
    <div className="flex justify-center items-center">
      <div className="w-72">
        <NativeSelectRoot>
          <NativeSelectField onChange={(e) => setSelector(Number(e.target.value))}>
            <option value="1">Очистить очередь(и перезапуск службы печати)</option>
            <option value="2">Выполнить тестовую печать</option>
            <option value="3">Логи lpf info</option>
            <option value="4">Логи lpf error</option>
            <option value="5">speedtest</option>
            <option value="6">speedtest --secure</option>
            <option value="7">Логи интерфейса сегодня</option>
            <option value="8">Логи интерфейса вчера</option>
            <option value="9">Логи интерфейса позавчера</option>
            <option value="10">Логи бота за 3 дня</option>
            <option value="11">Перезапустить бота</option>
            <option value="12">Логи перезагрузок за 3 дня</option>
            <option value="13">Тест сканера</option>
            <option value="14">Проверка подключения флешки</option>
            <option value="15">Исправления подключения флешки(применять с осторожностью) </option>
            <option value="16">Проверка видеороликов </option>
            <option value="17">Исправление прав видеороликов </option>
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
        if (selector === 7) {url = baseUrl + '/control/get-log-prod/1'}
        if (selector === 8) {url = baseUrl + '/control/get-log-prod/2'}
        if (selector === 9) {url = baseUrl + '/control/get-log-prod/3'}
        if (selector === 10) {url = baseUrl + '/control/get-log-bot'}
        if (selector === 11) {url = baseUrl + '/control/restart-bot'}
        if (selector === 12) {url = baseUrl + '/control/get-log-reboots'}
        if (selector === 13) {url = baseUrl + '/control/scanimage-test'}
        if (selector === 14) {url = baseUrl + '/control/check-usb-dir'}
        if (selector === 15) {url = baseUrl + '/control/correct-usb-dir'}
        if (selector === 16) {url = baseUrl + '/control/check-rek'}
        if (selector === 17) {url = baseUrl + '/control/correct-rek'}
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
