# BPFDemo
band pass filter demo

#

in this program, we will mix 200 HZ, 500HZ and 2000 HZ signal

then apply a 500 hz~1200hz band pass  to filter 500HZ out

#
	BPF := butter.NewBandPass2(wc, wd) // initialize 2nd order band pass butterworth filter
