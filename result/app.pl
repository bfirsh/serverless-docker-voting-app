use Mojolicious::Lite;
use Mojo::Pg;

helper pg => sub { state $pg = Mojo::Pg->new('postgresql://postgres@db/postgres') };

get '/votes' => sub {
  my $c = shift;

  my %data;
  my $results = $c->pg->db->query("SELECT vote, COUNT(id) AS count FROM votes GROUP BY vote");
  while (my $next = $results->hash) {
    $data{$next->{vote}} = $next->{count};
  }
  $c->render(json => {%data});
};

get '/' => sub {
  my $c = shift;
  $c->app->static->serve($c, 'index.html');
};

app->start;
