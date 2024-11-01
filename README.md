## Cypherlock

Ratchet based key expiry tool against forced decryption and for expiring
backups.

*PROOF OF CONCEPT CODE - DO NOT USE IN PRODUCTION UNLESS YOU KNOW EXACTLY WHAT YOU DO*

### Installation

```
go get -u -v github.com/cypherlock-pf/clv1/cmd/...
```

### Usage

First we create a new Cypherlock server:

```
$ cypherlockd -create
cypherlockd: minimal Cypherlock server
Server created.
SignatureKey: 8ad30073d3b5090eae94715304ec0916ea77bde2b3c3512e51ac55453bbe0c77

```

Then we let it run on the default interface (change interface with `-addr`):

```
$ cypherlockd -serve
cypherlockd: minimal Cypherlock server
Serving...
SignatureKey: 8ad30073d3b5090eae94715304ec0916ea77bde2b3c3512e51ac55453bbe0c77
```

Now we want to encrypt a time-locked `secret` file:

```
$ exec 3<secret; cypherlock -create -sigkey 8ad30073d3b5090eae94715304ec0916ea77bde2b3c3512e51ac55453bbe0c77
Please enter passphrase (no echo):
Please repeat passphrase (no echo):

Lock created. From "Wed Sep 19 22:40:27 +0000 UTC 2022" to "Wed Sep 19 23:10:27 +0000 UTC 2022"
```

To unlock the time-locked secret via the Cypherlock server and store it in file `secret2`:
```
$ exec 3>secret2; cypherlock -unlock -sigkey 8ad30073d3b5090eae94715304ec0916ea77bde2b3c3512e51ac55453bbe0c77
Please enter passphrase (no echo):
```

Now we have the content of the original `secret` file in `secret2`.
