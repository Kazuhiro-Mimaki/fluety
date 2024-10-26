# Fluety

A tool to monitor logs from your systems, applications that you use in your local environment.

# Quick start

### 1. Clone

```
git clone https://github.com/Kazuhiro-Mimaki/fluety.git
```

### 2. Pipe logs

Assuming the application that is running with npm, please move to your app's directory and execute start command.

```
npm run dev | {relative path to the fluety' build file}
```

For example, your app's path and fluety's path is same, "relative path to the fluety" is "../fluety/main".
