FROM suborbital/atmo:v0.4.1

COPY ./runnables.wasm.zip .

ENTRYPOINT [ "atmo" ]
