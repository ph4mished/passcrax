<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
  <body>
    <h1>PassCrax</h1>
    <p><em>A lightweight password-cracking tool for educational purposes</em></p>

    <h2>Overview</h2>
    <p>PassCrax is a personal project I developed in my early days of learning penetration testing. While more robust tools exist in the password-cracking space, I built this to deeply understand how hash cracking works. It started as a simple wordlist-based cracker and evolved to include additional features over time.</p>

    <h2>Features</h2>
    <div class="feature">
        <strong>Hash Identification</strong> - Detects hash types using regex pattern matching.
    </div>
    <div class="feature">
        <strong>Automatic Wordlist Selection</strong> - Scans all wordlists in the <code>/Wordlist</code> directory without manual selection.
    </div>
    <div class="feature">
        <strong>Multiple Attack Modes</strong> - Supports both <strong>wordlist attacks</strong> and <strong>brute-force</strong> methods.
    </div>

    <h2>⚠️ Legal Disclaimer</h2>
    <div class="disclaimer">
        <p><strong>WARNING: This tool is for LEGAL security testing only.</strong></p>
        
        <p>By using PassCrax, you agree to the following:</p>
        <ul>
            <li><strong>You must have explicit written permission</strong> to test any system or network.</li>
            <li><strong>Unauthorized hacking is illegal</strong> (violates laws like the CFAA, GDPR, Computer Misuse Act, etc.).</li>
            <li><strong>Use at your own risk</strong> - The developer is <strong>not</strong> liable for misuse or damages.</li>
            <li><strong>Intended for ethical hacking, pentesting, and educational research only.</strong></li>
        </ul>
        
        <blockquote><strong>"All risks are borne by the user. Misuse will result in legal consequences."</strong></blockquote>
    </div>

    <h2>Installation & Usage</h2>
    <p><em>git clone https://github.com/TAUREAN312/PassCrax.git
    cd PassCrax
    ruby PassCrax.rb</em></p>

    <h2>Contribution & License</h2>
    <ul>
        <li><strong>Open to feedback</strong> (Report issues or suggest improvements).</li>
        <li><strong>Educational use encouraged</strong> - Not for malicious purposes.</li>
    </ul>

    <h2>Why PassCrax?</h2>
    <p>This was a learning project—not meant to replace tools like Hashcat or John the Ripper. If you're exploring password security, feel free to test and contribute!</p>
</body>
</html>
