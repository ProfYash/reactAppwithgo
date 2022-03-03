import ReactDOM from 'react-dom'
import React from 'react'
import App from './App'
import {BrowserRouter} from 'react-router-dom'
import { CookiesProvider  } from 'react-cookie'
ReactDOM.render(<BrowserRouter><CookiesProvider> 
<App /></CookiesProvider> 
</BrowserRouter>,document.getElementById('root'))