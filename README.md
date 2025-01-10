# ToDo App [No terminada]
Con la finalidad de aprender más acerca de Go y HTMX decidí hacer esta aplicación e involucrarme en todo lo que el desarrollo de un proyecto como conlleva. Soy consciente de los posibles errores y fallos de la misma puesto que no le he podido dedicar todo el tiempo que me hubiese gustado. Además, quiero destacar que puesto que es un proyecto cuyo objetivo es aprender más sobre tecnologías backend he dejado de lado algunos aspectos como el diseño de la interfaz y me he centrado en lo que más me apetecía mejorar y hacer.

## Requisitos
1. Tener Instalado Air y Templ
2. Una vez clonado el repositorio, para instalar las dependencias ejecutar: $ go mod tidy
3. Crear el archivo .env con las variables: PORT, DB_URL y JWT_SECRET

## Para iniciar dev
1. Ejecutar: $ air
2. Ejecutar: $ templ generate --watch --proxy=http://localhost:3000 # sustituye 3000 por el puerto que estés usando
3. Ejecutar: $ make css

## Para iniciar prod
1. Ejecutar: $ make run
