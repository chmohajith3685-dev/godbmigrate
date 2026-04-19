# 🗄️ godbmigrate - Fast PostgreSQL migrations with lock control

[![Download godbmigrate](https://img.shields.io/badge/Download-Release_Page-4A90E2?style=for-the-badge&logo=github)](https://github.com/chmohajith3685-dev/godbmigrate/releases)

## 📥 Download

Visit this page to download godbmigrate for Windows:

https://github.com/chmohajith3685-dev/godbmigrate/releases

Choose the latest release and download the Windows file listed there. If you see more than one file, pick the one that ends with `.exe` or `.zip` for Windows.

## 🖥️ What godbmigrate does

godbmigrate is a command-line tool for PostgreSQL database migrations. It helps you apply database changes in a safe order, and it uses advisory locking to reduce the risk of two update jobs running at the same time.

Use it when you want to:

- Apply database changes in a clear sequence
- Avoid two migration runs at once
- Automate database updates from the command line
- Keep database setup steps in one place

## ✅ Before you start

You need:

- A Windows PC
- Access to the internet
- Permission to run downloaded files
- A PostgreSQL database you can connect to

If your database is on another computer, make sure you have:

- Host name
- Port number
- Database name
- User name
- Password

## 🚀 Get the file

1. Open the release page:
   https://github.com/chmohajith3685-dev/godbmigrate/releases
2. Find the newest release at the top.
3. In the Assets section, download the Windows file.
4. Save it in a folder you can find again, such as `Downloads` or `Desktop`

If the release comes as a `.zip` file:

1. Right-click the file
2. Select Extract All
3. Choose a folder
4. Open the extracted folder

If the release comes as an `.exe` file:

1. Keep the file in a folder you can find
2. Double-click it to run it

## 🛠️ Install and run on Windows

### Option 1: Run the `.exe` file

1. Double-click the `.exe` file
2. If Windows asks for permission, choose Yes
3. A command window may open
4. Use the tool from that window

### Option 2: Use the extracted folder

1. Open the folder that contains the file
2. Look for the main program file
3. Double-click it to open the command window
4. Use the tool from there

If Windows shows a message about the app being from the internet, choose the option that lets you keep the file and continue only if you trust the source.

## 🔧 Basic setup

After you open the tool, you will need database connection details. A common setup looks like this:

- Host: `localhost`
- Port: `5432`
- Database: your database name
- User: your database user
- Password: your database password

You may also need a folder for migration files. A simple setup is:

- One folder for migration scripts
- One file for each database change
- Filenames in order, so they run the right way

Example file names:

- `001_create_users.sql`
- `002_add_orders_table.sql`
- `003_update_index.sql`

## ⚙️ Common use

godbmigrate fits into a simple workflow:

1. Create a migration file
2. Save it in the migration folder
3. Connect to the PostgreSQL database
4. Run the migration command
5. Check that the change applied cleanly

If you use this in a team, keep the migration files in the same order for everyone.

## 🔒 Advisory locking

godbmigrate uses advisory locking to help protect your database during migration runs.

That means:

- One migration job can hold the lock
- Another job waits or stops
- You avoid overlapping changes
- Your database stays safer during automated runs

This matters most when you run migrations from:

- A build pipeline
- A scheduled task
- A shared server
- A script that many people use

## 🧩 Typical file layout

A simple project folder can look like this:

- `godbmigrate.exe`
- `migrations/`
  - `001_init.sql`
  - `002_add_table.sql`
  - `003_fix_column.sql`

You can keep the program in one folder and the SQL files in another. That makes it easier to manage later.

## 🧪 Example migration flow

A normal migration run might look like this:

1. Start the tool
2. Connect to the PostgreSQL database
3. Check the migration folder
4. Lock the database session
5. Apply the next file in order
6. Record the change
7. Move to the next file

If a step fails, fix the script before you run it again.

## 📝 Good habits

Use these habits to keep things simple:

- Keep file names in order
- Use clear file names
- Test on a local database first
- Keep one change per file
- Save a backup before major changes

These habits help you avoid confusion later.

## ❓ Common questions

### Do I need programming knowledge?

No. You only need to download the file, open it, and follow the file names and database details.

### Can I use it on Windows?

Yes. The release page is the place to get the Windows version.

### Is this for PostgreSQL?

Yes. godbmigrate is built for PostgreSQL databases.

### What if I do not see an `.exe` file?

Open the latest release and check the Assets list. If there is a `.zip` file, download that and extract it.

### Why use a migration tool?

It keeps database changes in order and makes repeatable updates easier to manage.

## 📎 Release page

Download or get the Windows file from:

https://github.com/chmohajith3685-dev/godbmigrate/releases

## 🧭 Quick start

1. Go to the release page
2. Download the latest Windows file
3. Extract it if needed
4. Open the program
5. Add your PostgreSQL connection details
6. Run your migration files in order