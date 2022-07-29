const crypto = require('crypto');

function base64URLEncode(str) {
	return str
		.toString('base64')
		.replace(/\+/g, '-')
		.replace(/\//g, '_')
		.replace(/=/g, '');
}

function sha256(buffer) {
	return crypto.createHash('sha256').update(buffer).digest();
}

const randomBytes = crypto.randomBytes(32);
console.log(randomBytes);
console.log(randomBytes.toString('base64'));
const code_verifier = base64URLEncode(randomBytes);
// const code_verifier = `fUap03sQDrci4ZqoFLa4jlqWhwAnJ6VLeXgilwsGKnc`;
const res = sha256(code_verifier);
console.log('sha:', res);
const code_challenge = base64URLEncode(res);
console.log('code_verifier: ' + code_verifier);
console.log('code_challenge: ' + code_challenge);
