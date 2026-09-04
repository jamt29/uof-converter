# uof-calculator

Conversor de unidades de medida en Go: temperatura (Celsius/Fahrenheit/Kelvin), longitud (metro/kilómetro/milla/pie) y peso (kilogramo/gramo/libra).

## Uso

```bash
go run . <categoria> <valor> <desde> <hasta>

# Ejemplos
go run . temp 100 C F     # 100.0000 C = 212.0000 F
go run . long 1 mi km   # 1.0000 mi = 1.6093 km
go run . peso 1 kg lb   # 1.0000 kg = 2.2046 lb
```

## Tests

```bash
go test ./... -v
```

## Prácticas de calidad integradas

### 1. Coding Standards (gofmt + golangci-lint)

**Qué práctica se aplicó**: se configuró `gofmt` que es una herramienta oficial integrada en el lenguaje Go, como formateador automático y `golangci-lint` (`.golangci.yml`) como linter estático, con reglas como `govet`, `errcheck`, `staticcheck`, `unused` y `revive` habilitadas. Ambas herramientas corren local (antes de crear un commit) y en CI (en cada push a rama main), así  ningún código que no cumpla con estandares  pasa desapercibido.

- **Qué problema evita**: inconsistencias de estilo entre distintos desarrolladores, errores  que  rompen en producción (variables sin usar, errores no chequeados, imports mal formateados) y discusiones  en code review que consumen tiempo sin agregar valor.
- **Cómo se relaciona**: reduce retrabajo. Un error de estilo o un `err` ignorado que es detectado por el linter en el momento en que se escribe el código cuesta un momento  en corregir; pero el mismo error descubierto en producción o en un code review tardío cuesta horas (o un incidente). Estandarizar el código automáticamente pone una herramienta que nunca se cansa ni se olvida de cumplir con ciertos aspectos.
- **Evidencia real**: al correr `golangci-lint run ./...` por primera vez sobre este mismo proyecto, encontró 5 issues reales de `revive` (constantes exportadas y paquetes sin comentario de documentación). Se corrigieron antes de subir el código — exactamente el tipo de detalle que, sin la herramienta, pasa a alguien que revisa o directo a `main`.

### 2. Integración Continua (CI) con GitHub Actions

**Qué práctica se aplicó**: se agregó un workflow (`.github/workflows/ci.yml`) que en cada `push` o `pull request` contra `main` corre automáticamente: verificación de formato (`gofmt -l`), lint (`golangci-lint`) y la suite de tests (`go test ./...`). Ningún cambio llega a `main` sin pasar por este control.

- **Qué problema evita**: la integración **"Big Bang"** — acumular varios cambios de varios desarrolladores (o de varias sesiones de trabajo) durante días o semanas y recién integrarlos todos juntos al final. Ese patrón produce conflictos de merge, bugs que son difíciles de rastrear porque hay demasiados cambios mezclados a la vez, y problemas de último momento justo antes de una entrega.
- **Cómo se relaciona**: en vez de integrar todo al final, cada cambio pequeño se valida automáticamente apenas se sube. Esto es integración continua en su forma más literal: integrar seguido, en piezas pequeñas, con feedback inmediato. Si algo se rompe, se sabe en minutos y sobre un diff pequeño — no semanas después sobre un diff enorme.
- **Cómo se va a demostrar**: el flujo de trabajo de este repo usa ramas de feature + Pull Request para cada cambio nuevo (en vez de subir todo directo a `main`). Cada PR dispara el mismo pipeline de CI: si el formateo, el lint o los tests fallan, el problema se detecta en el PR — antes de integrar — no después.

## Estructura

```
.
├── converter/           # Lógica de conversión (funciones puras, testeadas)
│   ├── temperature.go
│   ├── length.go
│   ├── weight.go
│   └── converter_test.go
├── main.go              # CLI
├── .golangci.yml        # Configuración del linter
└── .github/workflows/ci.yml  # Pipeline de CI
```
