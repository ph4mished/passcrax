<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
  <body>
    <h1>PassCrax</h1>
    <p><em>A lightweight and simple password-cracking tool for educational purposes</em></p>

<h2>Overview</h2>
    <p>Passcrax is a hash cracking tool which is built to be interactive, easier to use and beginner friendly. This tool expects little from the user and does much in the backend.
A little history of passcrax is that it was developed as a personal project and to help me to deeply understand how hash cracking word under the hood. It began as a simple wordlist-based hash cracker and evolved to include these additional features over time. 
Without wasting much time, let's dive into the <strong>features of passcrax</strong>;</p>

<h2>Features</h2>
    <div class="feature">
        <strong>Hash Identification</strong> - Detects hash types using regex pattern matching.
    </div>
    <div class="feature">
        <strong>Builtin Wordlist</strong> - PassCrax comes with its own wordlists files. This reduces the hassle of needing to download a wordlist file before using this tool. It scans all the wordlists in the <code>/Wordlist</code> directory without manual selection.
    </div>
        <div class="feature">
        <strong>User-Defined Wordlist Selection</strong> - PassCrax also supports external wordlists files. The only requirement from the user is to type "<em>load dictdir <<code>/path/to/wordlist_dir</code>></em>". <p>* This feature only support wordlist directories, not single files.</p>
    </div>
    <div class="feature">
        <strong>Hash & Hash File Support</strong> - Passcrax accepts both hashstring and hashfile inputs respectively. Launching this tool, the user is greeted with the status that shows it accepts hash inputs. Typing in the command "<em>load hashfile <<code>filepath</code>></em>"  (where 'filepath' is the path of your hash dump) makes the switch from hashstring status to hash file status.
    </div>
    <div class="feature">
        <strong>Rule Based Wordlist Mutation</strong> - Passcrax supports words mangling via rulefiles. This makes it possible to apply mutation rules to the wordlist, hence increasing the probability of getting your hashes cracked. This feature isn't available for file dump cracking yet.
    </div>
    <div class="feature">
        <strong>Multiple Cracking Modes</strong> - Supports three modes where the first two modes are <strong>wordlist cracking</strong> and <strong>brute-force</strong> methods. The third mode is kind of ad hoc which is <strong>auto</strong>. Auto mode is made to <em>smoothly transition from wordlist cracking to bruteforce if wordlist cracking was unsuccessful</em>.
    </div>

<h2>⚠️ Legal Disclaimer</h2>
    <div class="disclaimer">
        <p><strong>WARNING: This tool is for LEGAL security testing only.</strong></p>
        
<p>By using PassCrax, you agree to the following:</p>
        <ul>
            <li><strong>Only crack hashes you own.</strong></li>
            <li><strong>Unauthorized hacking is illegal</strong> (violates laws like the CFAA, GDPR, Computer Misuse Act, etc.).</li>
            <li><strong>Intended for ethical hacking, pentesting, and educational research only.</strong></li>
            <li><strong>Intended for ethical hacking, pentesting, CTF, and educational research only.</strong></li>
        </ul>
        
<blockquote><strong>"All risks are borne by the user. Misuse will result in legal consequences."</strong></blockquote>
    </div>

<p><div class="notice">
    <strong>👋 Hello guys</strong> - Please checkout my hash identifier <a href="https://github.com/ph4mished/hashpeek" target="_blank">"hashpeek"</a>.
</div></p>

    
<h2>Installation & Usage</h2>
    <p><pre><code><em>git clone https://github.com/ph4mished/passcrax.git<br>
    cd passcrax<br>
    go build passcrax.go<br>
    cd passcrax<br>
    go build <br>
    ./passcrax</em></code></pre></p>
<h2>Contribution & License</h2>
    <ul>
        <li><strong>Open to feedback</strong> (You can please report issues or suggest improvements).</li>
        <li><strong>Educational use is highly encouraged</strong></li>
    </ul>


<h2>Why PassCrax?</h2>
    <p>This was a learning project - to make hash cracking easier for ctf and beginners alike. If you're exploring password security, feel free to test and contribute!</p>
</body>
</html>
