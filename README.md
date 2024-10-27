# Fluety

A tool to monitor logs from your systems, applications that you use in your local environment.

<kbd><img width="1468" alt="image" src="https://github.com/user-attachments/assets/1131d594-2820-40ec-9ea9-8453e36254b6"></kbd>

# Quick start

### 1. Clone

```
git clone https://github.com/Kazuhiro-Mimaki/fluety.git
```

### 2. Pipe logs

Assuming the application that is running with npm, please move to your app's directory and execute start command.

```
npm run dev | {relative path to the fluety's build file}
```

For example, your app's path and fluety's path is same, relative path to the fluety is `../fluety/main`.

Let's access to http://localhost:8080.
