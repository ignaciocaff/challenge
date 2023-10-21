# Table of Contents

- [General Considerations](#general-considerations)``
- [Dockerization](#dockerization)
- [Bonus](#bonus)
- [Backend](#backend)
- [Bot](#bot)
- [Frontend](#frontend)

# General Considerations

### Important

Se decidió en el componente que renderiza finalmente los valores de las lineas utilizar
```javascript
.toLocaleString()?.replace(/[.,]/g, '')
```
Esto es porque en algunos casos el número que devolvía la linea era muy grande y como el requierimiento era que la api devuelva tipo number, se decidió hacerlo de esta manera para que se vea más amigable, eliminando los puntos y mostrando el número tal cual estaba originalmente en el .csv

Se configuro .gitattributes para que pise los archivos .csv el EOL a LF en lugar de CRLF porque sino rompe los tests al clonarlo en windows

- Frontend
    - Se utilizo props en un solo componente muy chico en lugar de context. En una relación directa ya que se considera que no había posibilidades de prop drilling
    - Se utilizo redux para el manejo de estado de la aplicación, en el caso de el estado "files", tendría más sentido si se removieran o adicionaran archivos de alguna manera
    - Cuando renderiza la aplicación se llama con useEffect a la consulta general. Luego en el filtro de búsqueda por nombre hay 2 variantes:
        - Si el input esta vacío y se ejecuta la consulta se muestra la lista completa (La busca d nuevo así en backend, no hay validacion de campo)
        - Si el input tiene un valor va al backend a la consulta por nombre
        - La consulta por nombre podría buscarse en el store de redux, pero se llama al backend para mostrar el uso del endpoint planteado en el punto adicional
    - Se utilizo tailwind para el manejo de estilos
    
- Backend:
    - Se configuro standardjs y el comando npm run lint
    - Se configuro el type = "module" en el package.json para poder utilizar import en lugar de require
    - Se utilizo express para el manejo de rutas y middlewares
    - Se utilizo axios para el manejo de peticiones externas
    - Se utilizo pino para el manejo de logs
    - Se utilizo mocha, chai y sinon para el manejo de pruebas unitarias y de api
    - Se utilizo dotenv para el manejo de variables de entorno (Se ejemplifica con .env y .env.prod) asumiento este último para la construcción en docker

# Dockerization

Se ha incluido un archivo Dockerfile en cada uno de los proyectos para facilitar la ejecución de los mismos en un entorno de desarrollo local. Para ejecutar la aplicación dockerizada se debe correr el siguiente comando en la raíz del proyecto:

```bash
docker-compose up --build -d
```
Lo que levantara los contendores de la api y el front en los puertos 8080 y 80 respectivamente.
Swagger estará disponible en /api-docs

# Bonus

Se puede probar la aplicación funcionando en el siguiente link -

Fue desplegada en una EC2 de aws.

# Backend

## Installation

1. Clona este repositorio.

2. Ejecuta los siguientes comandos
```bash
cd api
npm install
```

## Usage

Para iniciar la API, ejecuta:

```bash
cd api
npm start
```

La API estará disponible en http://localhost:8080
Swagger estará disponible en http://localhost:8080/api-docs

## Test

Se han incluido pruebas unitarias para asegurar la calidad del código disponibles en el archivo fileService.js y pruebas de api en el archivo api.js. Para ejecutar las pruebas, utiliza:

```bash
cd api
npm test
```

Esto ejecutará las pruebas utilizando Mocha, Chai y Sinon (Este último para mockear respuestas de la api externa que se consume archivos disponibles en /mocks).


## Structure

```bash
api
├── mocks                # Mocks para simular datos en el entorno de desarrollo (.csv).
├── src                  
│   ├── configs          # Configuraciones de variables de entorno.
│   ├── controllers      # Controladores para manejar las solicitudes.
│   ├── cors             # Configuración de CORS (Cross-Origin Resource Sharing).
│   ├── errors           # Manejo de errores personalizados (Middleware)
│   ├── logger           # Configuración y utilidades para el registro de eventos en consola usando pino.
│   ├── models           # Definición de modelos de negocio.
│   ├── routes           # Definición de las rutas de la API.
│   ├── services         # Lógica de negocio.
│   ├── thirdParty       # Integraciones con servicios de terceros.
├── test                 # Pruebas de la API (Contiene tests de valor limite y también de API).
├── .env                 # Archivo de variables de entorno (desarrollo).
├── .env.prod            # Archivo de variables de entorno (producción).
├── .nvmrc               # Versión de Node.js recomendada para la API.
├── Dockerfile           # Instrucciones para construir una imagen de Docker.
```


# Frontend

## Usage

The first thing you should do is install the npm dependencies for Angular.


### Docker




The front end will be available at  http://localhost


### Makefile

```bash
make install-app
```


The front end will be available at  http://localhost:4200

### Manual

```bash
cd cht-app
npm run i
```


The front end will be available at  http://localhost:4200


The first thing you should do is install the npm dependencies for Angular.

1. If you have Make installed, while in the project's root directory, execute the command 

```bash
make install-app
```

2. If you don't have Make installed, run the following commands:

```bash
cd cht-app
npm run i
```

Then you have several ways to run it

If you are using Docker, the front end will be available at http://localhost

If you don't have Docker, you have two alternatives:

1. If you have Make installed, while in the project's root directory, execute the command 

```bash
make run-chat-app
```

2. If you don't have Make installed, run the following commands:

```bash
cd cht-app
npm run dev
```

For both of the described cases, the front end will be available at  http://localhost:4200

## Structure

Below, I provide an incomplete but representative overview of the front-end structure

```bash
chat-app
├── config                 #  Configuration files for the application.                  
├── src     
│   ├── app      
│        ├── core          #  Core components that are used in the application.
│        ├── modules       #  Modules that are used in the application.
│        ├── routes        #  Routing configuration for the applicationn.
│   ├── assets 
│        ├── fonts              
│        ├── images              
│   ├── envinronments  
├── Dockerfile             # Instructions to build a Docker image.    
├── tailwind.config.js     # Tailwind configuration file with it's styles     
```
