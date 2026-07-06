import { useState } from 'react'
import axios from 'axios'
import './App.css'

// Базовый URL вашего сервера (замените localhost на ваш домен, если нужно)
const API_URL = 'http://0.0.0.0:8080/api/v1/auth' 

function App() {
  const [userName, setUserName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [message, setMessage] = useState('')
  const [loading, setLoading] = useState(false)

  const handleAuth = async (actionType) => {
    // Валидация заполнения полей в зависимости от действия
    if (!email || !password || (actionType === 'register' && !userName)) {
      setMessage('Пожалуйста, заполните все необходимые поля')
      return
    }

    setLoading(true)
    setMessage('')

    try {
      let response
      
      if (actionType === 'register') {
        // Регистрация: отправляем userName, email и password
        response = await axios.post(`${API_URL}/register`, { 
          userName, 
          email, 
          password 
        })
        setMessage('Успешная регистрация!')
      } else if (actionType === 'login') {
        // Вход: отправляем email и password
        response = await axios.post(`${API_URL}/login`, { email, password })
        setMessage(`Вы вошли! Токен: ${response.data.token || 'Успешно'}`)
      } else if (actionType === 'delete') {
        // Удаление: отправляем email и password в теле запроса (data)
        response = await axios.delete(`${API_URL}/delete_user`, { 
          data: { email, password } 
        })
        setMessage('Аккаунт успешно удален')
      }
    } catch (error) {
      setMessage(error.response?.data?.message || 'Произошла ошибка при запросе')
    } finally {
      setLoading(false)
    }
  }

  const isError = message.includes('ошибка') || message.includes('заполните')

  return (
    <>
      <section id="center">
        <div>
          <h1>Управление аккаунтом</h1>
        </div>

        <form onSubmit={(e) => e.preventDefault()}>
          <input
            type="text"
            placeholder="Имя пользователя (UserName)"
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
            disabled={loading}
          />
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            disabled={loading}
            required
          />
          <input
            type="password"
            placeholder="Пароль"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            disabled={loading}
            required
          />

          <div className="button-group">
            <button 
              type="button" 
              onClick={() => handleAuth('login')} 
              disabled={loading}
            >
              Войти
            </button>
            <button 
              type="button" 
              onClick={() => handleAuth('register')} 
              disabled={loading}
            >
              Регистрация
            </button>
          </div>

          <button 
            type="button" 
            className="danger-btn"
            onClick={() => handleAuth('delete')} 
            disabled={loading}
          >
            Удалить аккаунт
          </button>
        </form>

        {message && (
          <p className={`status-message ${isError ? 'error' : 'success'}`}>
            {message}
          </p>
        )}
      </section>

      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  )
}

export default App
