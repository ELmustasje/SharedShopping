# 🛒 Shared Shopping  

**Shared Shopping** is a collaborative shopping list web application where multiple users can create, manage, and share shopping lists in real-time. Designed for families, roommates, or groups of friends, it ensures everyone stays in sync when planning groceries.  

---

## ✨ Features  

- 👤 **User Authentication** — Secure login and signup with JWT-based authentication.  
- 📋 **Create & Manage Lists** — Create shopping lists and add/remove items with ease.  
- ✅ **Mark Items as Bought** — Bought items are automatically cleaned up after an hour.  
- 🔗 **Shareable Links** — List owners can generate invite links so others can join.  
- 📱 **Responsive UI** — Modern and mobile-friendly interface.  
- 🎨 **Polished UX** — Dropdown menus, modals, and a clean design for smooth collaboration.  

---

## 🏗️ Tech Stack  

- **Frontend:** React (TypeScript), Tailwind CSS, React Router, Axios  
- **Backend:** Go (Gin), JWT Authentication, Database (MySQL)  
- **Other:** Background jobs for automatic cleanup  

---

## 🚀 Getting Started  

### 1️⃣ Clone the repository  
```bash
git clone https://github.com/yourusername/shared-shopping.git
cd shared-shopping
```

### 2️⃣ Setup environment variables  
Create a `.env` file in the project root with your configuration:  

```env
JWT_SECRET=your-secret-key
DB_HOST=localhost
DB_USER=your-user
DB_PASS=your-password
DB_NAME=shared_shopping
```

### 3️⃣ Install & run backend  
```bash
cd backend
go mod tidy
go run main.go
```

### 4️⃣ Install & run portal
```bash
cd portal
npm install
npm run dev
```

---
