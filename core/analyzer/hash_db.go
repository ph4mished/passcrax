package analyzer

import "regexp"

type HASHALGO struct {
	Name     string
	PassCrax string
}

var HASH_DATABASE = map[*regexp.Regexp][]HASHALGO{
	//These hashtypes will be uncommented ones they get supported
	/*  regexp.MustCompile(`^(?i)[a-f0-9]{4}$`):  {
	    {Name: "CRC-16", PassCrax: "--"},
	    {Name: "CRC-16-CCITT", PassCrax: "--"},
	    {Name: "FCS-16", PassCrax: "--"},
	},*/

	/*regexp.MustCompile(`^(?i)[a-f0-9]{6}$`): {
	    {Name: "CRC-24", PassCrax: "--"},
	},*/

	regexp.MustCompile(`^(?i)[a-f0-9]{8}$`): {
		{Name: "Adler-32", PassCrax: "adler32"},
		{Name: "CRC-32", PassCrax: "crc32"},
		// { Name: "CRC-32B", PassCrax: "--" },
		{Name: "FNV-1-32", PassCrax: "fnv1-32"},
		{Name: "FNV-1a-32", PassCrax: "fnv1a-32"},
		// { Name: "Murmur3-32", PassCrax: "--" },
		/*  { Name: "FCS-32", PassCrax: "--" },
		    { Name: "GHash-32-3", PassCrax: "--" },
		    { Name: "GHash-32-5", PassCrax: "--" },
		    { Name: "Fletcher-32", PassCrax: "--" },
		    { Name:  "Joaat", PassCrax: "--" },*/
		//  { Name: "ELF-32", PassCrax: "--" },
		//  { Name: "XOR-32", PassCrax: "--" },
	},

	regexp.MustCompile(`^(?i)[a-f0-9]{16}$`): {
		//  { Name: "DES(Oracle)", PassCrax: "--" },
		// { Name: "LM", PassCrax: "--" },
		//{ Name: "MySQL323", PassCrax: "--" },
		{Name: "CRC-64", PassCrax: "crc64"},
		{Name: "FNV-1-64", PassCrax: "fnv1-64"},
		{Name: "FNV-1a-64", PassCrax: "fnv1a-64"},
		//  { Name: "Half-MD5", PassCrax: "--" },
	},

	regexp.MustCompile(`^(?i)[a-f0-9]{32}$`): {
		{Name: "MD5", PassCrax: "md5"},
		//  { Name: "NTLM", PassCrax: "--" },
		//{ Name: "LM", PassCrax: "--" },
		{Name: "MD4", PassCrax: "md4"},
		/*   {Name: "Double MD5", PassCrax: "--" },
		     { Name: "MD2", PassCrax: "--" },
		     { Name: "RIPEMD-128", PassCrax: "--" },
		     { Name: "BLAKE3-128", PassCrax: "--" },
		     { Name: "FNV-1-128", PassCrax: "--" },
		     { Name: "FNV-1a-128", PassCrax: "--" },
		     { Name: "Murmur3-128", PassCrax: "--" },
		     { Name: "SNEFRU-128", PassCrax: "--" },
		     { Name: "Skein256-128", PassCrax: "--" },
		     { Name: "Skein512-128", PassCrax: "--" },
		     { Name: "Tiger-128", PassCrax: "--" },
		     { Name: "Tiger128-3", PassCrax: "--" },
		     { Name: "Haval-128", PassCrax: "--" },*/
	},

	/* regexp.MustCompile(`^(?i)[a-f0-9]{34}$`): {
	    { Name: "CryptoCurrency(Adress)", PassCrax: "--" },
	},*/

	regexp.MustCompile(`^(?i)[a-f0-9]{40}$`): {
		{Name: "SHA-1", PassCrax: "sha1"},
		/*  { Name: "RIPEMD-160", PassCrax: "--" },
		    {Name: "Haval-160", PassCrax: "" },
		    { Name: "Double-SHA-1",PassCrax: "--" },
		    { Name: "Tiger-160", PassCrax: "--" },
		    { Name: "Tiger-160,3", PassCrax: "--" },*/
	},

	/*   regexp.MustCompile(`^(?i)[a-f0-9]{48}$`): {
	     { Name: "Haval-192", PassCrax: "--" },
	     { Name: "Tiger-192", PassCrax: "--" },
	     { Name: "Tiger-192-3", PassCrax: "--" },
	     { Name: "SHA-1(Oracle)", PassCrax: "--" },
	 },*/

	/*  regexp.MustCompile(`^(?i)[a-f0-9]{49}$`): {
	    { Name: "Citrix Netscaler", PassCrax: "--" },
	},*/

	/*  regexp.MustCompile(`^\$2[abxy]\$\d{2}\$[./A-Za-z0-9]{53}$`): {
	    { Name: "bcrypt", PassCrax: "bcrypt" },
	},*/

	regexp.MustCompile(`^(?i)[a-f0-9]{56}$`): {
		{Name: "SHA-224", PassCrax: "sha224"},
		{Name: "SHA3-224", PassCrax: "sha3-224"},
		/* { Name: "Haval-224", PassCrax: "--" },
		   { Name: "Skein-256(224)", PassCrax: "--" },
		   { Name:  "Skein-512(224)", PassCrax: "--" },
		   { Name: "Whirlpool-224", PassCrax: "--" },*/
	},

	regexp.MustCompile(`^(?i)[a-f0-9]{64}$`): {
		{Name: "SHA-256", PassCrax: "sha256"},
		{Name: "SHA3-256", PassCrax: "sha3-256"},
		/*  { Name: "BLAKE2s", PassCrax: "--" },
		    { Name: "BLAKE3-256", PassCrax: "--" },
		    { Name: "RIPEMD-256", PassCrax: "--" },
		    { Name: "Haval-256", PassCrax: "--" },
		    {Name: "Gost", PassCrax: "--" },
		    {Name: "GOST R 34.11-94", PassCrax: "--" },
		    { Name: "Gost-CryptoPro S-Box", PassCrax: "--" },
		    { Name: "SNEFRU-256", PassCrax: "--" },
		    { Name: "EDON-R-256", PassCrax: "--" },
		    { Name: "Skein256-256", PassCrax: "--" },
		    { Name: "Skein512-256", PassCrax: "--" },
		    { Name: "Whirlpool-256", PassCrax: "--" },*/
	},

	/* regexp.MustCompile(`^(?i)[a-f0-9]{80}$`): {
	    { Name: "RIPEMD-320", PassCrax: "--" },
	},*/

	regexp.MustCompile(`^(?i)[a-f0-9]{96}$`): {
		{Name: "SHA-384", PassCrax: "sha384"},
		{Name: "SHA3-384", PassCrax: "sha3-384"},
		/*   { Name: "Skein512-384", PassCrax: "--" },
		     { Name: "Skein1024-384", PassCrax: "--" },
		     { Name: "Whirlpool-384", PassCrax: "--" },*/
	},

	/*  regexp.MustCompile(`^(?i)[a-f0-9]{102}$`): {
	    { Name: "Skein1024-408", PassCrax: "--" },
	},*/

	regexp.MustCompile(`^(?i)[a-f0-9]{128}$`): {
		{Name: "SHA-512", PassCrax: "sha512"},
		{Name: "SHA3-512", PassCrax: "sha3-512"},
		/* { Name: "BLAKE2b", PassCrax: "--" },
		   { Name: "BLAKE3-512", PassCrax: "--" },
		   {Name: "Salsa10", PassCrax: "--" },
		   {Name: "Salsa20", PassCrax: "--" },
		   { Name: "Skein512-512", PassCrax: "--" },
		   { Name: "Whirlpool", PassCrax: "--" },
		   { Name: "EDON-R-512", PassCrax: "--" },
		   { Name: "Whirlpool-1", PassCrax: "--" },
		   { Name: "Whirlpool-2", PassCrax: "--" },*/
	},

	/*  regexp.MustCompile(`^(?i)[a-f0-9]{256}$`): {
	    { Name: "Skein512-1024", PassCrax: "--" },
	    { Name: "Skein1024-1024", PassCrax: "--" },
	},*/
}
