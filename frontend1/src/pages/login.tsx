import React, { useState, useCallback } from 'react';
import {
  AppBar,
  Toolbar,
  Typography,
  Container,
  Box,
  Card,
  CardContent,
  CardActions,
  TextField,
  Button,
  Alert,
  Divider,
  InputAdornment,
  IconButton,
  CssBaseline,
} from '@mui/material';
import {
  AccountCircle,
  Lock,
  Visibility,
  VisibilityOff,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import axios from 'axios'; 

// ธีมสีหลัก: สีเทาอ่อน (#E0E0E0)
const PRIMARY_COLOR = '#E0E0E0';
const API_URL = 'http://localhost:8080/login'; 

// อินเทอร์เฟซสำหรับข้อมูลฟอร์ม
interface LoginForm {
  email: string;
  password: string;
}

// Interface สำหรับข้อมูลที่คาดว่าจะได้รับจาก Backend (data ภายใน response.data)
interface LoginData {
    token: string;
    role: 'pregnant' | 'doctor' | 'admin' | string;
    name: string;
    id: number | string;
}

// Validation Rules
const validateEmail = (email: string): string => {
  return email ? '' : 'กรุณากรอกอีเมลหรือชื่อผู้ใช้';
};

const validatePassword = (password: string): string => {
  if (!password) return 'กรุณากรอกรหัสผ่าน';
  if (password.length < 6) return 'รหัสผ่านต้องมีอย่างน้อย 6 ตัวอักษร';
  return '';
};


const LoginView: React.FC = () => {
  const navigate = useNavigate();

  // State Management
  const [formData, setFormData] = useState<LoginForm>({
    email: '',
    password: '',
  });
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [loginError, setLoginError] = useState<string | null>(null);
  const [emailError, setEmailError] = useState<string>('');
  const [passwordError, setPasswordError] = useState<string>('');

  // Handlers
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
    // เคลียร์ Error เมื่อมีการพิมพ์
    if (e.target.name === 'email') setEmailError('');
    if (e.target.name === 'password') setPasswordError('');
  };

  const validateForm = useCallback((): boolean => {
    const emailValidation = validateEmail(formData.email);
    const passwordValidation = validatePassword(formData.password);

    setEmailError(emailValidation);
    setPasswordError(passwordValidation);

    return !emailValidation && !passwordValidation;
  }, [formData]);

  const navigateUser = useCallback((role: string) => {
    switch (role.toLowerCase()) {
      case 'pregnant':
        navigate('/home'); 
        break;
      case 'doctor':
        navigate('/home/doctor-dashboard'); 
        break;
      case 'admin':
        navigate('/admin-dashboard'); 
        break;
      default:
        navigate('/'); 
    }
  }, [navigate]);

  const login = useCallback(async () => {
    if (!validateForm()) return;

    setLoading(true);
    setLoginError(null);

    try {
      const response = await axios.post(API_URL, {
        username: formData.email, // Backend รับ 'username'
        password: formData.password,
      });

      // 🎯 การเข้าถึงข้อมูลโดยตรงโดยใช้ 'any' และ Assert เป็น LoginData
      const apiResponseData = response.data as any;
      const loginData = apiResponseData.data as LoginData; // เข้าถึง property 'data' และ Assert Type เป็น LoginData
      
      const { token, role, name, id } = loginData;

      // จัดเก็บ Token และข้อมูลผู้ใช้
      localStorage.setItem('authToken', token);
      localStorage.setItem('userRole', role);
      localStorage.setItem('userName', name);
      localStorage.setItem('userID', String(id)); 

      // นำทางผู้ใช้
      navigateUser(role);

    } catch (error) {
      // 🎯 ใช้ Type Assertion 'as any' สำหรับจัดการ Error Response
      console.error('Login Error:', (error as any).response ? (error as any).response.data : (error as any).message);

      let errorMessage = 'เกิดข้อผิดพลาดในการเข้าสู่ระบบ';
      const responseData = (error as any).response?.data;

      if (responseData && responseData.error) {
        errorMessage = responseData.error; 
      } else if ((error as any).message && (error as any).message.includes('Network Error')) {
        errorMessage = 'ไม่สามารถเชื่อมต่อกับ Server ได้ กรุณาตรวจสอบสถานะ Backend';
      }

      setLoginError(errorMessage);
    } finally {
      setLoading(false);
    }
  }, [formData, navigateUser, validateForm]);

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <CssBaseline />
      
      {/* App Bar (เทียบเท่า v-app-bar) */}
      <AppBar position="fixed" sx={{ bgcolor: PRIMARY_COLOR, boxShadow: 'none' }}>
        <Toolbar>
          <Typography variant="h6" component="div" fontWeight="medium">
            Maternal and Child Health Handbook
          </Typography>
        </Toolbar>
      </AppBar>

      {/* Main Content (เทียบเท่า v-main) */}
      <Box 
        component="main" 
        sx={{ 
          flexGrow: 1, 
          pt: 8, 
          display: 'flex', 
          alignItems: 'center', 
          justifyContent: 'center',
          bgcolor: 'grey.50' 
        }}
      >
        <Container component="div" maxWidth="sm">
          <Card sx={{ p: 3, m: 1, elevation: 6, borderRadius: '8px' }}>
            <CardContent>
              <Typography 
                variant="h5" 
                component="div" 
                fontWeight="bold" 
                textAlign="center" 
                mb={3}
              >
                เข้าสู่ระบบ
              </Typography>

              {/* Form (เทียบเท่า v-form) */}
              <Box component="form" onSubmit={(e) => { e.preventDefault(); login(); }}>
                
                {/* Email Field */}
                <TextField
                  fullWidth
                  name="email"
                  label="อีเมล/ชื่อผู้ใช้"
                  type="text"
                  value={formData.email}
                  onChange={handleChange}
                  error={!!emailError}
                  helperText={emailError}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <AccountCircle />
                      </InputAdornment>
                    ),
                  }}
                />

                {/* Password Field */}
                <TextField
                  fullWidth
                  name="password"
                  label="รหัสผ่าน"
                  value={formData.password}
                  onChange={handleChange}
                  error={!!passwordError}
                  helperText={passwordError}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 4 }}
                  type={showPassword ? 'text' : 'password'}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <Lock />
                      </InputAdornment>
                    ),
                    endAdornment: (
                      <InputAdornment position="end">
                        <IconButton
                          onClick={() => setShowPassword(!showPassword)}
                          edge="end"
                          size="small"
                        >
                          {showPassword ? <VisibilityOff /> : <Visibility />}
                        </IconButton>
                      </InputAdornment>
                    ),
                  }}
                />

                {/* Submit Button */}
                <Button
                  type="submit"
                  fullWidth
                  variant="contained"
                  sx={{ mt: 1, bgcolor: PRIMARY_COLOR, color: 'black' }}
                  disabled={loading}
                  onClick={login}
                >
                  {loading ? 'กำลังเข้าสู่ระบบ...' : 'เข้าสู่ระบบ'}
                </Button>
              </Box>

              {/* Error Alert */}
              {loginError && (
                <Alert severity="error" onClose={() => setLoginError(null)} sx={{ mt: 2 }}>
                  {loginError}
                </Alert>
              )}
            </CardContent>

            {/* Actions (เทียบเท่า v-card-actions) */}
            <CardActions sx={{ display: 'flex', flexDirection: 'column', pt: 0, pb: 3 }}>
              
              {/* Forget Password */}
              <Button 
                onClick={() => navigate('/forget.password')} 
                size="small" 
                sx={{ mb: 1, color: 'text.secondary' }}
              >
                ลืมรหัสผ่าน?
              </Button>
              
              <Divider sx={{ width: '100%', my: 1 }} />
              
              {/* Register Button */}
              <Button 
                onClick={() => navigate('/register')}
                variant="outlined" 
                sx={{ mt: 1, borderColor: 'grey.400', color: 'text.primary' }}
              >
                สมัครสมาชิก
              </Button>
            </CardActions>
          </Card>
        </Container>
      </Box>
    </Box>
  );
};

export default LoginView;