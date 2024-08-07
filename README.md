# Convertidor de unidades en Go

¡Bienvenido al proyecto de Convertidor de Unidades en Go! Este repositorio acompaña al post "Nivel 6: Creemos un Convertidor de Unidades en Go" de nuestro blog gognition.pro donde aprendimos a crear un convertidor de unidades multifuncional.

## 📋 Descripción

Aprende a crear un Convertidor de Unidades en Go. Explora interfaces, mapas y manejo de errores

## 🚀 Comenzando

Para utilizar este proyecto, asegúrate de tener Go instalado en tu sistema.

### Prerrequisitos

- Go (versión 1.15 o superior recomendada)

### Instalación

1. Haz un fork de este repositorio haciendo clic en el botón "Fork" en la parte superior derecha de esta página.

2. Clona tu fork a tu máquina local:

```bash
git clone https://github.com/TU_USUARIO/gognition-nivel6-convertidor-de-unidades.git
```

## 💻️ Uso

Para ejecutar el convertidor de unidades:
```bash
go run main.go
```

El programa mostrará salidas como:
```
===== Convertidor de Unidades =====
1. Longitud
2. Peso
3. Temperatura
4. Salir
Seleccione el tipo de conversión (1-4): 2
Ingrese el valor a convertir: 50
Ingrese la unidad de origen: kg
Ingrese la unidad de destino: lb
Resultado: 50.00 kg = 110.23 lb
```

## 🏗️ Compilación (Build)

Para compilar el proyecto y crear un ejecutable, puedes usar los siguientes comandos dependiendo de tu sistema operativo objetivo:
```bash
go build -o convertidor-unidades
```

Para Windows (desde Linux o macOS):
```bash
GOOS=windows GOARCH=amd64 go build -o convertidor-unidades.exe
```

Para macOS (desde Windows o Linux):
```bash
GOOS=darwin GOARCH=amd64 go build -o convertidor-unidades
```

Para Linux (desde Windows o macOS):
```bash
GOOS=linux GOARCH=amd64 go build -o convertidor-unidades
```

Después de la compilación, puedes ejecutar el programa usando:
```
./convertidor-unidades # En Linux o macOS
convertidor-unidades.exe  # En Windows
```

## Visita gognition.pro https://www.gognition.pro/