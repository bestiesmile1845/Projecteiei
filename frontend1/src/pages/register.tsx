import React, { useState, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
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
  Divider,
  InputAdornment,
  IconButton,
  CssBaseline,
  Alert,
} from '@mui/material';
import {
  Email,
  Lock,
  LockOpen,
  Visibility,
  VisibilityOff,
} from '@mui/icons-material';

// ธีมสีหลัก: สีเทาอ่อน (#E0E0E0)
const PRIMARY_COLOR = '#E0E0E0';
const API_URL = 'http://localhost:3000/register'; // สมมติ API

// Interface สำหรับข้อมูลฟอร์ม
interface RegisterForm {
  email: string;
  password: string;
  confirmPassword: string;
}

// Interface สำหรับ Error State
interface FormErrors {
  email: string;
  password: string;
  confirmPassword: string;
}

// --- Validation Functions ---

const validateEmail = (email: string): string => {
  if (!email) return 'กรุณากรอกอีเมล';
  const emailRegex = /.+@.+\..+/;
  if (!emailRegex.test(email)) return 'รูปแบบอีเมลไม่ถูกต้อง';
  return '';
};

const validatePassword = (password: string): string => {
  if (!password) return 'กรุณากรอกรหัสผ่าน';
  if (password.length < 6) return 'รหัสผ่านต้องมีอย่างน้อย 6 ตัวอักษร';
  return '';
};

const validateConfirmPassword = (
  confirmPassword: string,
  password: string
): string => {
  if (!confirmPassword) return 'กรุณายืนยันรหัสผ่าน';
  if (confirmPassword !== password) return 'รหัสผ่านไม่ตรงกัน';
  return '';
};

const RegisterView: React.FC = () => {
  const navigate = useNavigate();

  // State Management
  const [formData, setFormData] = useState<RegisterForm>({
    email: '',
    password: '',
    confirmPassword: '',
  });
  const [errors, setErrors] = useState<FormErrors>({
    email: '',
    password: '',
    confirmPassword: '',
  });
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [registerError, setRegisterError] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);

  // Handlers
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
    // เคลียร์ Error เมื่อมีการพิมพ์
    setErrors({ ...errors, [name]: '' });
    setRegisterError(null);
    setSuccessMessage(null);
  };

  const validateForm = useCallback((): boolean => {
    const emailError = validateEmail(formData.email);
    const passwordError = validatePassword(formData.password);
    const confirmPasswordError = validateConfirmPassword(
      formData.confirmPassword,
      formData.password
    );

    const newErrors = {
      email: emailError,
      password: passwordError,
      confirmPassword: confirmPasswordError,
    };

    setErrors(newErrors);

    return !emailError && !passwordError && !confirmPasswordError;
  }, [formData]);

  const register = useCallback(async () => {
    if (!validateForm()) return;

    setLoading(true);
    setRegisterError(null);
    setSuccessMessage(null);

    // --- จำลองการเรียก API Register (TODO: แทนที่ด้วย axios/fetch จริง) ---
    try {
      // **TODO: นำเข้าและใช้ axios.post(API_URL, { email: formData.email, password: formData.password, ... }) ที่นี่**
      
      console.log('Attempting to register with:', formData);

      // จำลองการหน่วงเวลาการเรียก API 2 วินาที
      await new Promise(resolve => setTimeout(resolve, 2000)); 

      // สมมติว่า Register สำเร็จ
      setSuccessMessage('สมัครสมาชิกสำเร็จ! กำลังนำทางไปหน้าเข้าสู่ระบบ...');
      
      // นำทางไปหน้า Login หลังจาก 2 วินาที
      setTimeout(() => {
        navigate('/login');
      }, 2000);

    } catch (error) {
      console.error('Registration Error:', error);
      // **TODO: จัดการ error response จาก API จริง**
      setRegisterError('เกิดข้อผิดพลาดในการสมัครสมาชิก (โปรดตรวจสอบ Backend)');
    } finally {
      setLoading(false);
    }
  }, [formData, navigate, validateForm]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    register();
  };

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <CssBaseline />
      
      {/* App Bar (เทียบเท่า v-app-bar) */}
      <AppBar position="fixed" flat sx={{ bgcolor: 'white', borderBottom: 1, borderColor: 'grey.300', boxShadow: 'none' }}>
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
          bgcolor: 'grey.50',
          p: 2
        }}
      >
        <Container component="div" maxWidth="sm">
          {/* Register Card (เทียบเท่า v-card) */}
          <Card sx={{ p: 3, m: 1, elevation: 6, borderRadius: '8px', maxWidth: '500px' }}>
            <CardContent>
              <Typography 
                variant="h5" 
                component="div" 
                fontWeight="bold" 
                textAlign="center" 
                mb={3}
              >
                สมัครสมาชิก
              </Typography>

              {/* Form (เทียบเท่า v-form) */}
              <Box component="form" onSubmit={handleSubmit} noValidate>
                
                {/* Email Field */}
                <TextField
                  fullWidth
                  name="email"
                  label="อีเมล"
                  type="email"
                  value={formData.email}
                  onChange={handleChange}
                  error={!!errors.email}
                  helperText={errors.email}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <Email />
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
                  error={!!errors.password}
                  helperText={errors.password}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
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

                {/* Confirm Password Field */}
                <TextField
                  fullWidth
                  name="confirmPassword"
                  label="ยืนยันรหัสผ่าน"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                  error={!!errors.confirmPassword}
                  helperText={errors.confirmPassword}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 4 }}
                  type={showPassword ? 'text' : 'password'}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <LockOpen />
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

                {/* Register Button */}
                <Button
                  type="submit"
                  fullWidth
                  variant="contained"
                  sx={{ mt: 2, bgcolor: PRIMARY_COLOR, color: 'black' }}
                  disabled={loading}
                >
                  {loading ? 'กำลังสมัครสมาชิก...' : 'สมัครสมาชิก'}
                </Button>
              </Box>

              {/* Success/Error Alert */}
              {successMessage && (
                <Alert severity="success" sx={{ mt: 2 }}>
                  {successMessage}
                </Alert>
              )}
              {registerError && (
                <Alert severity="error" onClose={() => setRegisterError(null)} sx={{ mt: 2 }}>
                  {registerError}
                </Alert>
              )}

            </CardContent>

            {/* Actions (กลับสู่หน้าเข้าสู่ระบบ) */}
            <CardActions sx={{ display: 'flex', flexDirection: 'column', pt: 0, pb: 3 }}>
              
              <Divider sx={{ width: '100%', my: 2 }} />
              
              <Button 
                onClick={() => navigate('/login')}
                variant="outlined" 
                size="small"
                sx={{ mt: 1, borderColor: 'grey.400', color: 'text.primary' }}
              >
                กลับสู่หน้าเข้าสู่ระบบ
              </Button>
            </CardActions>
          </Card>
        </Container>
      </Box>
    </Box>
  );
};

export default RegisterView;