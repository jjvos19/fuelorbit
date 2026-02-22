CREATE TABLE adm_empresa (
    id_empresa INT IDENTITY(1,1) PRIMARY KEY,
    nombre NVARCHAR(200) NOT NULL,
    nit NVARCHAR(50),
    direccion NVARCHAR(250),
    telefono NVARCHAR(50),
    email NVARCHAR(150),
    wallet_address NVARCHAR(150) NOT NULL UNIQUE,
    hash_empresa VARCHAR(500),
    fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
    estado NVARCHAR(50) NOT NULL DEFAULT 'ACTIVA'
);


---------------------------------

CREATE TABLE adm_cisterna (
    id_cisterna INT IDENTITY(1,1) PRIMARY KEY,
    id_empresa INT NOT NULL,
    placa NVARCHAR(20) NOT NULL UNIQUE,
    capacidad_maxima DECIMAL(10,2) NOT NULL,
    modelo_sensor NVARCHAR(100),
    estado NVARCHAR(50) NOT NULL DEFAULT 'ACTIVA',
    fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),

    CONSTRAINT FK_cisterna_empresa FOREIGN KEY (id_empresa)
        REFERENCES adm_empresa(id_empresa)
);

--------------------------

CREATE TABLE adm_conductor (
    id_conductor INT IDENTITY(1,1) PRIMARY KEY,
    id_empresa INT NOT NULL,
    nombre NVARCHAR(150) NOT NULL,
    licencia NVARCHAR(50) NOT NULL UNIQUE,
    estado NVARCHAR(50) NOT NULL DEFAULT 'ACTIVO',
    hash_identidad VARCHAR(500),
    fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),

    CONSTRAINT FK_conductor_empresa FOREIGN KEY (id_empresa)
        REFERENCES adm_empresa(id_empresa)
);
---------------------------------------
CREATE TABLE adm_surtidor (
    id_surtidor INT IDENTITY(1,1) PRIMARY KEY,
    id_empresa INT NOT NULL,
    nombre NVARCHAR(150) NOT NULL,
    direccion NVARCHAR(250),
    estado NVARCHAR(50) NOT NULL DEFAULT 'ACTIVO',
    fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),

    CONSTRAINT FK_surtidor_empresa FOREIGN KEY (id_empresa)
        REFERENCES adm_empresa(id_empresa)
);

-----------------------------------
CREATE TABLE adm_cisterna_conductor (
    id_relacion INT IDENTITY(1,1) PRIMARY KEY,
    id_cisterna INT NOT NULL,
    id_conductor INT NOT NULL,
    fecha_inicio DATETIME2 NOT NULL,
    fecha_fin DATETIME2 NULL,

    CONSTRAINT FK_cc_cisterna FOREIGN KEY (id_cisterna)
        REFERENCES adm_cisterna(id_cisterna),

    CONSTRAINT FK_cc_conductor FOREIGN KEY (id_conductor)
        REFERENCES adm_conductor(id_conductor)
);
-------------------
CREATE TABLE trk_data (
    id_trk_data BIGINT IDENTITY(1,1) PRIMARY KEY,
    id_cisterna INT NOT NULL,
    latitude DECIMAL(10,7) NOT NULL,
    longitude DECIMAL(10,7) NOT NULL,
    volumen_actual DECIMAL(10,2) NOT NULL,
    temperatura DECIMAL(5,2),
    estado NVARCHAR(20),
    nivel_combustible_sensor DECIMAL(10,2),
    hash_device VARCHAR(500),
    hash_bck VARCHAR(500),

    CONSTRAINT FK_trk_cisterna FOREIGN KEY (id_cisterna)
        REFERENCES adm_cisterna(id_cisterna)
);
---------------
CREATE TABLE trk_entrega (
    id_entrega BIGINT IDENTITY(1,1) PRIMARY KEY,
    id_cisterna INT NOT NULL,
    id_surtidor INT NOT NULL,
    volumen_entregado DECIMAL(10,2) NOT NULL,
    tokens_quemados DECIMAL(10,2) NOT NULL,
    hash_transaccion VARCHAR(500),
    fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),

    CONSTRAINT FK_entrega_cisterna FOREIGN KEY (id_cisterna)
        REFERENCES adm_cisterna(id_cisterna),

    CONSTRAINT FK_entrega_surtidor FOREIGN KEY (id_surtidor)
        REFERENCES adm_surtidor(id_surtidor)
);

