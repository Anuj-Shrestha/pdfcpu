/*
Copyright 2018 The pdfcpu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

const (
	usage = `pdfcpu is a tool for PDF manipulation written in Go. 
	
Usage:
	
	pdfcpu command [arguments]
		
The commands are:
	
	validate	validate PDF against PDF 32000-1:2008 (PDF 1.7)
	optimize	optimize PDF by getting rid of redundant page resources
	split		split multi-page PDF into several single-page PDFs
	merge		concatenate 2 or more PDFs
	extract		extract images, fonts, content, pages, metadata
	trim		create trimmed version
	attach		list, add, remove, extract embedded file attachments
	perm		list, add user access permissions
	encrypt		set password protection		
	decrypt		remove password protection
	changeupw	change user password
	changeopw	change owner password
	stamp		add stamps
	watermark	add watermarks
	version		print version
   
	Single-letter Unix-style supported for commands and flags.

Use "pdfcpu help [command]" for more information about a command.`

	usageValidate     = "usage: pdfcpu validate [-v(erbose)|vv] [-mode strict|relaxed] [-upw userpw] [-opw ownerpw] inFile"
	usageLongValidate = `Validate checks inFile for specification compliance.

verbose, v ... turn on logging
        vv ... verbose logging
      mode ... validation mode
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
		
The validation modes are:

 strict ... (default) validates against PDF 32000-1:2008 (PDF 1.7)
relaxed ... like strict but doesn't complain about common seen spec violations.`

	usageOptimize     = "usage: pdfcpu optimize [-v(erbose)|vv] [-stats csvFile] [-upw userpw] [-opw ownerpw] inFile [outFile]"
	usageLongOptimize = `Optimize reads inFile, removes redundant page resources like embedded fonts and images and writes the result to outFile.

verbose, v ... turn on logging
        vv ... verbose logging
     stats ... appends a stats line to a csv file with information about the usage of root and page entries.
               useful for batch optimization and debugging PDFs.
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
   outFile ... output pdf file (default: inFile-new.pdf)`

	usageSplit     = "usage: pdfcpu split [-v(erbose)|vv] [-upw userpw] [-opw ownerpw] inFile outDir"
	usageLongSplit = `Split generates a set of single page PDFs for the input file in outDir.

verbose, v ... turn on logging
        vv ... verbose logging
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
    outDir ... output directory`

	usageMerge     = "usage: pdfcpu merge [-v(erbose)|vv] outFile inFile..."
	usageLongMerge = `Merge concatenates a sequence of PDFs/inFiles to outFile.

verbose, v ... turn on logging
        vv ... verbose logging
   outFile ... output pdf file
   inFiles ... a list of at least 2 pdf files subject to concatenation.`

	usageExtract     = "usage: pdfcpu extract [-v(erbose)|vv] -mode image|font|content|page|meta [-pages pageSelection] [-upw userpw] [-opw ownerpw] inFile outDir"
	usageLongExtract = `Extract exports inFile's images, fonts, content or pages into outDir.

verbose, v ... turn on logging
        vv ... verbose logging
      mode ... extraction mode
     pages ... page selection
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
    outDir ... output directory

 The extraction modes are:

  image ... extract images (supported PDF filters: Flate, DCTDecode, JPXDecode)
   font ... extract font files (supported font types: TrueType)
content ... extract raw page content
   page ... extract single page PDFs
   meta ... extract all metadata (page selection does not apply)`

	usageTrim     = "usage: pdfcpu trim [-v(erbose)|vv] [-pages pageSelection] [-upw userpw] [-opw ownerpw] inFile [outFile]"
	usageLongTrim = `Trim generates a trimmed version of inFile for selected pages.

verbose, v ... turn on logging
        vv ... verbose logging
     pages ... page selection
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
   outFile ... output pdf file (default: inFile-new.pdf)`

	usagePageSelection = `<pages> selects pages for processing and is a comma separated list of expressions:

	Valid expressions are:

	even ... include even pages           odd ... include odd pages
  	   # ... include page #               #-# ... include page range
 	  !# ... exclude page #              !#-# ... exclude page range
 	  n# ... exclude page #              n#-# ... exclude page range

	  #- ... include page # - last page    -# ... include first page - page #
 	 !#- ... exclude page # - last page   !-# ... exclude first page - page #
 	 n#- ... exclude page # - last page   n-# ... exclude first page - page #

	n serves as an alternative for !, since ! needs to be escaped with single quotes on the cmd line.

e.g. -3,5,7- or 4-7,!6 or 1-,!5 or odd,n1`

	usageAttachList    = "pdfcpu attach list [-v(erbose)|vv] [-upw userpw] [-opw ownerpw] inFile"
	usageAttachAdd     = "pdfcpu attach add [-v(erbose)|vv] [-upw userpw] [-opw ownerpw] inFile file..."
	usageAttachRemove  = "pdfcpu attach remove [-v(erbose)|vv] [-upw userpw] [-opw ownerpw] inFile [file...]"
	usageAttachExtract = "pdfcpu attach extract [-v(erbose)|vv] [-upw userpw] [-opw ownerpw] inFile outDir [file...]"

	usageAttach = "usage: " + usageAttachList +
		"\n       " + usageAttachAdd +
		"\n       " + usageAttachRemove +
		"\n       " + usageAttachExtract

	usageLongAttach = `Attach manages embedded file attachments.
	
verbose, v ... turn on logging
        vv ... verbose logging
      perm ... user access permissions
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
    outDir ... output directory`

	usagePermList = "pdfcpu perm list [-v(erbose)|vv] [-upw userpw] [-opw ownerpw] inFile"
	usagePermAdd  = "pdfcpu perm add [-v(erbose)|vv] [-perm none|all] [-upw userpw] -opw ownerpw inFile"

	usagePerm = "usage: " + usagePermList +
		"\n       " + usagePermAdd

	usageLongPerm = `Perm manages user access permissions.
	
verbose, v ... turn on logging
        vv ... verbose logging
      perm ... user access permissions
       upw ... user password
       opw ... owner password
    sinFile ... input pdf file`

	usageEncrypt     = "usage: pdfcpu encrypt [-v(erbose)|vv] [-mode rc4|aes] [-key 40|128] [perm none|all] [-upw userpw] [-opw ownerpw] inFile [outFile]"
	usageLongEncrypt = `Encrypt sets a password protection based on user and owner password.

verbose, v ... turn on logging
        vv ... verbose logging
      mode ... algorithm (default=aes)
       key ... key length in bits (default=128)
      perm ... user access permissions
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
   outFile ... output pdf file`

	usageDecrypt     = "usage: pdfcpu decrypt [-v(erbose)|vv] [-upw userpw] [-opw ownerpw] inFile [outFile]"
	usageLongDecrypt = `Decrypt removes a password protection.

verbose, v ... turn on logging
        vv ... verbose logging
       upw ... user password
       opw ... owner password
    inFile ... input pdf file
   outFile ... output pdf file`

	usageChangeUserPW     = "usage: pdfcpu changeupw [-v(erbose)|vv] [-opw ownerpw] inFile upwOld upwNew"
	usageLongChangeUserPW = `Changeupw changes the user password.
	
verbose, v ... turn on logging
        vv ... verbose logging
       opw ... owner password, required unless = ""
    inFile ... input pdf file
    upwOld ... old user password
    upwNew ... new user password`

	usageChangeOwnerPW     = "usage: pdfcpu changeopw [-v(erbose)|vv] [-upw userpw] inFile opwOld opwNew"
	usageLongChangeOwnerPW = `Changeopw changes the owner password.
	
verbose, v ... turn on logging
        vv ... verbose logging
       upw ... user password, required unless = ""
    inFile ... input pdf file
    opwOld ... old owner password (supply user password on initial changeopw)
    opwNew ... new owner password`

	usageWMDescription = `<description> is a comma separated configuration string containing:
	
    1st entry: the display string
               or an image file name with one the of extensions '.png', '.tif' or '.tiff' 
               or a PDF file name with extension .pdf followed by an optional page number (default=1) separated by ':'

    optional entries:

         (defaults: 'f:Helvetica, p:24, s:0.5 rel, c:0.5 0.5 0.5, r:0, d:1, o:1, m:0')

      f: fontname, a basefont, supported are: Helvetica, Times-Roman, Courier
      p: fontsize in points, in combination with absolute scaling only.
      s: scale factor, 0.0 <= x <= 1.0 followed by optional 'abs|rel' or 'a|r'.
      c: color: 3 fill color intensities, where 0.0 < i < 1.0, eg 1.0, 0.0 0.0 = red (default:0.5 0.5 0.5 = gray)
      r: rotation, where -180.0 <= x <= 180.0
      d: render along diagonal, 1..lower left to upper right, 2..upper left to lower right (if present overrules r!)
      o: opacity, where 0.0 <= x <= 1.0
      m: render mode: 0 ... fill
                      1 ... stroke
                      2 ... fill & stroke

    Only one of rotation and diagonal is allowed.

e.g. 'Draft'                                                  'logo.png'
     'Draft, d:2'                                             'logo.tif, o:0.5, s:0.5 abs, r:0'
     'Intentionally left blank, p:48'                         'some.pdf, r:45' 
     'Confidental, f:Courier, s:0.75, c: 0.5 0.0 0.0, r:20'   'some.pdf:3, r:-90, s:0.75'`

	usageStamp     = "usage: pdfcpu stamp [-v(erbose)|vv] [-pages pageSelection] description inFile [outFile]"
	usageLongStamp = `Stamp adds stamps for selected pages. 

 verbose, v ... turn on logging
         vv ... verbose logging
      pages ... page selection
description ... font, font size, text, color, image/pdf file name, pdf page#, rotation, opacity, scale factor, render mode
     inFile ... input pdf file
    outFile ... output pdf file (default: inFile-new.pdf)

` + usageWMDescription

	usageWatermark     = "usage: pdfcpu watermark [-v(erbose)|vv] [-pages pageSelection] description inFile [outFile]"
	usageLongWatermark = `Watermark adds watermarks for selected pages. 

 verbose, v ... turn on logging
         vv ... verbose logging
      pages ... page selection
description ... font, font size, text, color, image/pdf file name, pdf page#, rotation, opacity, scale factor, render mode
     inFile ... input pdf file
    outFile ... output pdf file (default: inFile-new.pdf)

` + usageWMDescription

	usageVersion     = "usage: pdfcpu version"
	usageLongVersion = "Version prints the pdfcpu version"
)
