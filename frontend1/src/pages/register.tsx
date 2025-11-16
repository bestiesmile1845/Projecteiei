import React, { useState, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
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
  Person,
  Phone,
  Cake,
  Badge,
  VpnKey,
} from '@mui/icons-material';

// ธีมสีหลัก: สีเทาอ่อน (#E0E0E0)
const PRIMARY_COLOR = '#E0E0E0';
const API_URL = 'http://localhost:8080'; // Endpoint สำหรับ CreatePregnantWoman

// Interface สำหรับข้อมูลฟอร์ม
interface RegisterForm {
  fullName: string;
  phoneNumber: string;
  age: string;
  hn: string;
  citizenID: string;
  email: string;
  password: string;
  confirmPassword: string;
}

// Interface สำหรับ Error State
interface FormErrors {
  fullName: string;
  phoneNumber: string;
  age: string;
  hn: string;
  citizenID: string;
  email: string;
  password: string;
  confirmPassword: string;
}

// --- Validation Functions ---

const validateFullName = (name: string): string => {
  return name ? '' : 'กรุณากรอกชื่อ-นามสกุล';
};

const validatePhoneNumber = (phone: string): string => {
  if (!phone) return 'กรุณากรอกเบอร์โทรศัพท์';
  if (!/^\d{10}$/.test(phone)) return 'รูปแบบเบอร์โทรศัพท์ไม่ถูกต้อง (10 หลัก)';
  return '';
};

const validateAge = (age: string): string => {
  if (!age) return 'กรุณากรอกอายุ';
  const ageNum = parseInt(age);
  if (isNaN(ageNum) || ageNum < 15 || ageNum > 50) return 'อายุต้องอยู่ระหว่าง 15-50 ปี';
  return '';
};

const validateCitizenID = (id: string): string => {
  if (!id) return 'กรุณากรอกเลขบัตรประชาชน';
  if (!/^\d{13}$/.test(id)) return 'เลขบัตรประชาชนต้องมี 13 หลัก';
  return '';
};

const validateHN = (hn: string): string => {
  return hn ? '' : 'กรุณากรอก HN';
};

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
    fullName: '',
    phoneNumber: '',
    age: '',
    hn: '',
    citizenID: '',
    email: '',
    password: '',
    confirmPassword: '',
  });
  const [errors, setErrors] = useState<FormErrors>({
    fullName: '',
    phoneNumber: '',
    age: '',
    hn: '',
    citizenID: '',
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
    const fullNameError = validateFullName(formData.fullName);
    const phoneNumberError = validatePhoneNumber(formData.phoneNumber);
    const ageError = validateAge(formData.age);
    const hnError = validateHN(formData.hn);
    const citizenIDError = validateCitizenID(formData.citizenID);
    const emailError = validateEmail(formData.email);
    const passwordError = validatePassword(formData.password);
    const confirmPasswordError = validateConfirmPassword(
      formData.confirmPassword,
      formData.password
    );

    const newErrors = {
      fullName: fullNameError,
      phoneNumber: phoneNumberError,
      age: ageError,
      hn: hnError,
      citizenID: citizenIDError,
      email: emailError,
      password: passwordError,
      confirmPassword: confirmPasswordError,
    };

    setErrors(newErrors);

    return (
      !fullNameError &&
      !phoneNumberError &&
      !ageError &&
      !hnError &&
      !citizenIDError &&
      !emailError &&
      !passwordError &&
      !confirmPasswordError
    );
  }, [formData]);

  const register = useCallback(async () => {
    if (!validateForm()) return;

    setLoading(true);
    setRegisterError(null);
    setSuccessMessage(null);

    try {
      const response = await axios.post(API_URL, {
        FullName: formData.fullName,
        Email: formData.email,
        Password: formData.password,
        Username: formData.email, // สมมติให้ใช้ Email เป็น Username
        PhoneNumber: formData.phoneNumber,
        Age: parseInt(formData.age), // แปลงเป็นตัวเลขตาม Go struct
        HN: formData.hn,
        CitizenID: formData.citizenID,
      });

      console.log('Registration Success:', response.data);

      setSuccessMessage('สมัครสมาชิกสำเร็จ! กำลังนำทางไปหน้าเข้าสู่ระบบ...');
      
      setTimeout(() => {
        navigate('/login');
      }, 2000);

    } catch (error) {
      const responseData = (error as any).response?.data;
      let errorMessage = 'เกิดข้อผิดพลาดในการสมัครสมาชิก';

      if (responseData && responseData.error) {
        errorMessage = responseData.error; // ดึง error จาก Backend (เช่น username already exists)
      } else if ((error as any).message && (error as any).message.includes('Network Error')) {
        errorMessage = 'ไม่สามารถเชื่อมต่อกับ Server ได้ กรุณาตรวจสอบสถานะ Backend';
      }
      setRegisterError(errorMessage);
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
      
      {/* App Bar */}
      <AppBar position="fixed" sx={{ bgcolor: 'white', borderBottom: 1, borderColor: 'grey.300', boxShadow: 'none' }}>        
        <Toolbar>
          <Typography variant="h6" component="div" fontWeight="medium">
            Maternal and Child Health Handbook
          </Typography>
        </Toolbar>
      </AppBar>

      {/* Main Content */}
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
          {/* Register Card */}
          <Card sx={{ p: 3, m: 1, elevation: 6, borderRadius: '8px', maxWidth: '500px' }}>
            <CardContent>
              <Typography 
                variant="h5" 
                component="div" 
                fontWeight="bold" 
                textAlign="center" 
                mb={3}
              >
                สมัครสมาชิก (สตรีมีครรภ์)
              </Typography>

              <Box component="form" onSubmit={handleSubmit} noValidate>
                
                {/* Full Name Field */}
                <TextField
                  fullWidth
                  name="fullName"
                  label="ชื่อ-นามสกุล"
                  value={formData.fullName}
                  onChange={handleChange}
                  error={!!errors.fullName}
                  helperText={errors.fullName}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
                  InputProps={{
                    startAdornment: (<InputAdornment position="start"><Person /></InputAdornment>),
                  }}
                />

                {/* Phone Number Field */}
                <TextField
                  fullWidth
                  name="phoneNumber"
                  label="เบอร์โทรศัพท์"
                  value={formData.phoneNumber}
                  onChange={handleChange}
                  error={!!errors.phoneNumber}
                  helperText={errors.phoneNumber}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
                  InputProps={{
                    startAdornment: (<InputAdornment position="start"><Phone /></InputAdornment>),
                  }}
                />

                {/* Age Field */}
                <TextField
                  fullWidth
                  name="age"
                  label="อายุ"
                  type="number"
                  value={formData.age}
                  onChange={handleChange}
                  error={!!errors.age}
                  helperText={errors.age}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
                  InputProps={{
                    startAdornment: (<InputAdornment position="start"><Cake /></InputAdornment>),
                  }}
                />

                {/* HN Field */}
                <TextField
                  fullWidth
                  name="hn"
                  label="HN (Hospital Number)"
                  value={formData.hn}
                  onChange={handleChange}
                  error={!!errors.hn}
                  helperText={errors.hn}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
                  InputProps={{
                    startAdornment: (<InputAdornment position="start"><VpnKey /></InputAdornment>),
                  }}
                />

                {/* Citizen ID Field */}
                <TextField
                  fullWidth
                  name="citizenID"
                  label="เลขบัตรประชาชน"
                  value={formData.citizenID}
                  onChange={handleChange}
                  error={!!errors.citizenID}
                  helperText={errors.citizenID}
                  variant="outlined"
                  size="small"
                  required
                  sx={{ mb: 2 }}
                  InputProps={{
                    startAdornment: (<InputAdornment position="start"><Badge /></InputAdornment>),
                  }}
                />

                {/* Email Field */}
                <TextField
                  fullWidth
                  name="email"
                  label="อีเมล (ใช้เป็นชื่อผู้ใช้)"
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
                    startAdornment: (<InputAdornment position="start"><Email /></InputAdornment>),
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
                    startAdornment: (<InputAdornment position="start"><Lock /></InputAdornment>),
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
                    startAdornment: (<InputAdornment position="start"><LockOpen /></InputAdornment>),
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