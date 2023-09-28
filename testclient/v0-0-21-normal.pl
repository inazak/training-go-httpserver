#!/usr/bin/perl
use strict;
use warnings;
use LWP;
use JSON;

## 1. create user `beatles` and `carpenters`
## 2. login `beatles`
## 3. add two task
## 4. get task and print
## 5. login `carpenters`
## 6. add two task
## 7. get task and print


my ($host, $res, $token, $tasks) = ("http://localhost:18080", undef, undef, undef);

$res = request(
	'GET', $host, '/health',
	undef,
	undef,
	200)
	or die;

$res = request(
	'POST', $host, '/register',
	undef,
	{name => "beatles", password => "123"},
	200)
  or die;

$res = request(
	'POST', $host, '/register',
	undef,
	{name => "carpenters", password => "456"},
	200)
  or die;

$res = request(
	'POST', $host, '/login',
	undef,
	{name => "beatles", password => "123"},
	200)
  or die;

$token = decode_json($res->content)->{"accesstoken"};

$res = request(
	'POST', $host, '/task',
	undef,
	undef,
	401) #unauthorized
  or die;

$res = request(
	'POST', $host, '/task',
	{Authorization => "Bearer " . $token},
	{title => "she loves you"},
	200)
  or die;

$res = request(
	'POST', $host, '/task',
	{Authorization => "Bearer " . $token},
	{title => "a day in the life"},
	200)
  or die;

$res = request(
	'GET', $host, '/task',
	{ Authorization => "Bearer " . $token },
	undef,
	200)
  or die;

$tasks = decode_json($res->content);
print "beatles", "\n";
print "title list:", "\n";
for my $task (@$tasks) {
	print "- ", $task->{"title"}, "\n";
}

$res = request(
	'POST', $host, '/login',
	undef,
	{name => "carpenters", password => "456"},
	200)
  or die;

$token = decode_json($res->content)->{"accesstoken"};

$res = request(
	'POST', $host, '/task',
	{Authorization => "Bearer " . $token},
	{title => "goodbye to love"},
	200)
  or die;

$res = request(
	'POST', $host, '/task',
	{Authorization => "Bearer " . $token},
	{title => "top of the world"},
	200)
  or die;

$res = request(
	'GET', $host, '/task',
	{ Authorization => "Bearer " . $token },
	undef,
	200)
  or die;

$tasks = decode_json($res->content);
print "carpenters", "\n";
print "title list:", "\n";
for my $task (@$tasks) {
	print "- ", $task->{"title"}, "\n";
}

exit 0;

sub request {
	my $method = shift;
	my $host   = shift;
	my $path   = shift;
	my $header = shift;
	my $data   = shift;
	my $expect = shift;

	my $ua = LWP::UserAgent->new();
	my $req = HTTP::Request->new($method => $host . $path);

	if (defined $header) {
		for my $k (keys %$header) {
			$req->header($k => $header->{$k});
		}
	}
	if (defined $data) {
		$req->content(encode_json($data));
	}

	my $res = $ua->request($req);

	if ($res->code == $expect) {
		print "ok ", $res->code, " ", $path, "\n";
	} else {
		print "Fail ", $res->code, " ", $path, " ", $res->content, "\n";
		return undef;
	}

	return $res;
}

