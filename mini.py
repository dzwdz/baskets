# bird's eye view of the crypto involved

from hashlib import blake2s, scrypt, sha256, pbkdf2_hmac

def emulatedDevice(site):
	# If this was a real device, it would show the name of the site and wait
	# for authentication. This ensures that even if you use a compromised
	# computer to log into an account, only that account will be endangered.
	# The screen prevents a MiTM attack too.
	print(f"[device] requested key for {site}")

	# The device holds a secret generated as in BIP39. This step would be
	# usually performed by the setup software on a trusted computer.
	mnemonic = "ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic"
	mnemonic_pass = "TREZOR"
	secret = pbkdf2_hmac(
		"sha512",
		mnemonic.encode("utf-8"), ("mnemonic" + mnemonic_pass).encode("utf-8"),
		2048
	)[:32] # 32 is blake2s's maximum key size

	return blake2s(site.encode("utf-8"), key=secret).digest()

def getBaseKey(passphrase, n, r, p):
	salt = emulatedDevice("baskets")
	return scrypt(passphrase.encode("utf-8"), salt=salt, n=n, r=r, p=p, dklen=32)

def getSiteKey(bk, service, counter):
	# Site keys are generated using the user's passphrase, so even if the
	# device gets lost or stolen, the passwords are still somewhat safe.
	#
	# The counter lets one create multiple passwords for the same service
	# name. It should only be used to rotate passwords; different accounts
	# on the same site should use a different service string.
	toHash = bk + emulatedDevice(service) + str(counter).encode("utf-8")
	return sha256(toHash).digest()

bk = getBaseKey("correct horse battery staple", 1024, 8, 1)
print(f"base key = {bk.hex()[:8]}...")
print(f"dzwdz@lichess site key = {getSiteKey(bk, 'dzwdz@lichess', 0).hex()[:8]}...")
print(f"lobste.rs site key = {getSiteKey(bk, 'lobste.rs', 0).hex()[:8]}...")
