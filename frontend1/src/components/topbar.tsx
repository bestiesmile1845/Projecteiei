import React, { useState, useCallback } from 'react';
import {
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Divider,
  Box,
  Avatar,
  Button,
  Container,
  CssBaseline,
} from '@mui/material';
import {
  Menu as MenuIcon,
  Logout as LogoutIcon,
  Dashboard as DashboardIcon,
  MedicalInformation as StethoscopeIcon, 
  Vaccines as SyringeIcon,
  FaceRetouchingNatural as BabyFaceIcon,
  PeopleAlt as FamilyIcon,
  PregnantWoman as PregnantIcon,
  Book as BookIcon,
  Settings as SettingsIcon,
  AccountCircle as AccountIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom'; // สมมติว่าใช้ React Router

// ธีมสีหลัก: สีเทาอ่อน (#E0E0E0)
const PRIMARY_COLOR = '#E0E0E0';
const PRIMARY_COLOR_LIGHT = 'grey.100'; // สำหรับ sheet/sheet
const BLACK_TEXT_COLOR = 'text.primary'; // สำหรับข้อความสีดำ

// ข้อมูลเมนูลิงก์
const links = [
  { icon: DashboardIcon, text: 'หน้าหลัก Dashboard', route: '/home' },
  { icon: StethoscopeIcon, text: 'บันทึกการตรวจ', route: '/home/checkup' },
  { icon: SyringeIcon, text: 'ประวัติวัคซีน', route: '/home/vaccines' },
  { icon: BabyFaceIcon, text: 'ประวัติตั้งครรภ์ปัจจุบัน', route: '/home/pregnancy/current' },
  { icon: FamilyIcon, text: 'ประวัติสุขภาพและครอบครัว', route: '/home/health/family' },
  { icon: PregnantIcon, text: 'นับลูกดิ้น', route: '/home/kick-counter' },
  { icon: BookIcon, text: 'ข้อมูลสุขภาพ', route: '/home/info' },
  { icon: SettingsIcon, text: 'ตั้งค่าบัญชี', route: '/home/settings' },
];

interface MainLayoutProps {
  children: React.ReactNode;
}

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  const [drawerOpen, setDrawerOpen] = useState(false);
  const navigate = useNavigate(); // ใช้ React Router Hook

  // ข้อมูลผู้ใช้จำลอง
  const userName = 'คุณแม่ A';
  const userEmail = 'user@example.com';

  const toggleDrawer = useCallback(() => {
    setDrawerOpen((prev) => !prev);
  }, []);

  const handleLogout = useCallback(() => {
    console.log('Logging out and navigating to login page...');
    // นำทางกลับไปหน้า login (สมมติว่า route name คือ '/login')
    navigate('/login');
  }, [navigate]);

  const handleLinkClick = (route: string) => {
    navigate(route);
    setDrawerOpen(false);
  };

  const drawerContent = (
    <Box sx={{ width: 250 }} role="presentation">
      {/* User Sheet (เทียบเท่า v-sheet) */}
      <Box sx={{ p: 2, bgcolor: PRIMARY_COLOR_LIGHT }}>
        <Avatar sx={{ mb: 1, bgcolor: PRIMARY_COLOR, color: BLACK_TEXT_COLOR, width: 64, height: 64 }}>
          <AccountIcon sx={{ fontSize: 30 }} />
        </Avatar>
        <Typography variant="subtitle1" fontWeight="bold">
          {userName}
        </Typography>
        <Typography variant="caption" color="text.secondary">
          {userEmail}
        </Typography>
      </Box>

      <Divider />

      {/* Navigation List (เทียบเท่า v-list nav) */}
      <List component="nav">
        {links.map((item) => (
          <ListItem key={item.route} disablePadding>
            <ListItemButton 
              onClick={() => handleLinkClick(item.route)}
              // ใน React Router เราจะใช้ useLocation hook เพื่อตรวจสอบ active
              // แต่เพื่อความง่ายในตัวอย่างนี้ เราจะละไว้
            >
              <ListItemIcon sx={{ minWidth: 40 }}>
                <item.icon />
              </ListItemIcon>
              <ListItemText primary={item.text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>

      {/* Logout Button (เทียบเท่า v-slot:append) */}
      <Box sx={{ p: 2, position: 'absolute', bottom: 0, width: '100%' }}>
        <Button 
          fullWidth 
          variant="contained" 
          color="inherit" 
          sx={{ bgcolor: 'grey.800', color: 'white' }}
          onClick={handleLogout}
        >
          <LogoutIcon sx={{ mr: 1 }} />
          ออกจากระบบ
        </Button>
      </Box>
    </Box>
  );

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      
      {/* App Bar (เทียบเท่า v-app-bar) */}
      <AppBar 
        position="fixed" 
        sx={{ bgcolor: PRIMARY_COLOR, boxShadow: 'none' }}
      >
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            edge="start"
            onClick={toggleDrawer}
            sx={{ color: BLACK_TEXT_COLOR, mr: 2 }}
          >
            <MenuIcon />
          </IconButton>
          <Typography 
            variant="h6" 
            noWrap 
            component="div" 
            fontWeight="medium"
            sx={{ color: BLACK_TEXT_COLOR }}
          >
            Maternal and Child Health Handbook
          </Typography>
        </Toolbar>
      </AppBar>

      {/* Navigation Drawer (เทียบเท่า v-navigation-drawer) */}
      <Drawer
        variant="temporary"
        open={drawerOpen}
        onClose={toggleDrawer}
        ModalProps={{
          keepMounted: true, // For mobile responsiveness
        }}
        sx={{
          '& .MuiDrawer-paper': { boxSizing: 'border-box', width: 250 },
        }}
      >
        {drawerContent}
      </Drawer>

      {/* Main Content (เทียบเท่า v-main) */}
      <Box
        component="main"
        sx={{ 
          flexGrow: 1, 
          p: 0, // Padding ถูกย้ายไปที่ Container ด้านล่าง
          width: '100%',
          mt: 8, // ชดเชยความสูงของ App Bar
        }}
      >
        <Container 
          maxWidth={false} // ทำให้ Container กว้างสุดเท่าที่ Box จะรับได้
          sx={{ 
            py: 4, // 2 เท่าของ py-8 (Vuetify units)
            px: 2, 
            minHeight: 'calc(100vh - 64px)', // min-hight
            // Vuetify's max-width/max-height is complex, using standard CSS for simplicity
          }}
        >
          {children}
        </Container>
      </Box>
    </Box>
  );
};

export default MainLayout;