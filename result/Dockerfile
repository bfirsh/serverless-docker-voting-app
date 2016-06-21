FROM perl:5
RUN curl -L https://cpanmin.us | perl - -M https://cpan.metacpan.org -n Mojolicious Mojo::Pg
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp
CMD [ "perl", "./app.pl" , "cgi"]
