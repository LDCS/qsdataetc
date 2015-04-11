// qsdataetc collects info from various sources on a linux box - mainly etc files, but also some commands, and outputs it into csv-formatted file
//
// This enables an enterprise csvfile-based ETL environment to monitor linux servers and desktops
package main

import (
	"fmt"
	"github.com/LDCS/genutil"
	"github.com/LDCS/qslinux/etcfstab"
	"github.com/LDCS/qslinux/etchosts"
	"github.com/LDCS/qslinux/etcservice"
	"github.com/LDCS/qslinux/etcshadow"
	"github.com/LDCS/qslinux/etcuser"
	"github.com/LDCS/sflag"
)

var (
	opt = struct {
		Usage    string "wrapper for etc files functionality"
		Filelist string "List of system files to csv (hosts, user)	| hosts,user,fstab,service"
		Odir     string "output directory (default is to use stdout) 	|"
		Verbose  bool   "verbosity | false"
	}{}
)

func doHosts(_mybox string) {
	ostr := ""
	ostrVerbose := ""
	hostsmap := etchosts.DoListHostsdata(opt.Verbose)
	if true {
		ostr += fmt.Sprintf("box,%s\n", etchosts.Header())
		strs := genutil.SortedUniqueKeys(etchosts.Keys_String2PtrHostsdata(&hostsmap))
		strsDone := map[string]bool{}
		for _, kk := range strs {
			strsDone[kk] = true
			hosts, _ := hostsmap[kk]
			ostr += fmt.Sprintf("%s,%s\n", _mybox, hosts.Csv())
		}
	} else {
		for _, kk := range etchosts.SortedKeys_String2PtrHostsdata(&hostsmap) {
			elem := hostsmap[kk]
			ostr += elem.Sprint()
		}
	}
	ofile := ""
	switch {
	case opt.Odir == "":
		ofile = "/dev/stdout"
	default:
		ofile = opt.Odir + "/qsdataetc.hosts." + _mybox + ".csv"
	}
	genutil.WriteStringToFile(ostrVerbose+ostr, ofile)
}

func doFstab(_mybox string) {
	ostr := ""
	ostrVerbose := ""
	fstabmap := etcfstab.DoListFstabdata(opt.Verbose)
	if true {
		ostr += fmt.Sprintf("box,%s\n", etcfstab.Header())
		strs := genutil.SortedUniqueKeys(etcfstab.Keys_String2PtrFstabdata(&fstabmap))
		strsDone := map[string]bool{}
		for _, kk := range strs {
			strsDone[kk] = true
			fstab, _ := fstabmap[kk]
			ostr += fmt.Sprintf("%s,%s\n", _mybox, fstab.Csv())
		}
	} else {
		for _, kk := range etcfstab.SortedKeys_String2PtrFstabdata(&fstabmap) {
			elem := fstabmap[kk]
			ostr += elem.Sprint()
		}
	}
	ofile := ""
	switch {
	case opt.Odir == "":
		ofile = "/dev/stdout"
	default:
		ofile = opt.Odir + "/qsdataetc.fstab." + _mybox + ".csv"
	}
	genutil.WriteStringToFile(ostrVerbose+ostr, ofile)
}

func doService(_mybox string) {
	ostr := ""
	ostrVerbose := ""
	servicemap := etcservice.DoListServicedata(opt.Verbose)
	if true {
		ostr += fmt.Sprintf("box,%s\n", etcservice.Header())
		strs := genutil.SortedUniqueKeys(etcservice.Keys_String2PtrServicedata(&servicemap))
		strsDone := map[string]bool{}
		for _, kk := range strs {
			strsDone[kk] = true
			service, _ := servicemap[kk]
			ostr += fmt.Sprintf("%s,%s\n", _mybox, service.Csv())
		}
	} else {
		for _, kk := range etcservice.SortedKeys_String2PtrServicedata(&servicemap) {
			elem := servicemap[kk]
			ostr += elem.Sprint()
		}
	}
	ofile := ""
	switch {
	case opt.Odir == "":
		ofile = "/dev/stdout"
	default:
		ofile = opt.Odir + "/qsdataetc.service." + _mybox + ".csv"
	}
	genutil.WriteStringToFile(ostrVerbose+ostr, ofile)
}

func doUser(_mybox string) {
	ostr := ""
	ostrVerbose := ""
	usermap := etcuser.DoListUserdata(opt.Verbose)
	shadowmap := etcshadow.DoListShadowdata(opt.Verbose)
	if true {
		ostr += fmt.Sprintf("box,%s,%s\n", etcuser.Header(), etcshadow.Header())
		strs := genutil.SortedUniqueKeys(etcuser.Keys_String2PtrUserdata(&usermap))
		strsDone := map[string]bool{}
		for _, kk := range strs {
			strsDone[kk] = true
			user, _ := usermap[kk]
			shadow, _ := shadowmap[kk]
			ostr += fmt.Sprintf("%s,%s,%s\n", _mybox, user.Csv(), shadow.Csv())
		}
	} else {
		for _, kk := range etcuser.SortedKeys_String2PtrUserdata(&usermap) {
			elem := usermap[kk]
			ostr += elem.Sprint()
		}
	}
	ofile := ""
	switch {
	case opt.Odir == "":
		ofile = "/dev/stdout"
	default:
		ofile = opt.Odir + "/qsdataetc.user." + _mybox + ".csv"
	}
	genutil.WriteStringToFile(ostrVerbose+ostr, ofile)
}

func main() {
	mybox := genutil.Hostname()
	sflag.Parse(&opt)
	if opt.Verbose {
		fmt.Println("\nStarting on ", mybox, "verbose=", opt.Verbose)
	}

	for _, filetype := range genutil.AnySplit(opt.Filelist, ",") {
		switch filetype {
		case "hosts":
			doHosts(mybox)
		case "fstab":
			doFstab(mybox)
		case "service":
			doService(mybox)
		case "user":
			doUser(mybox)
		}
	}
}
