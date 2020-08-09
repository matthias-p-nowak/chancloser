/*
This package "chancloser" closes a channel when the last sender of that channel finishes work.
For this to work, a sender has to claim a channel with "chancloser.Claim(ch)" and using "defer chancloser.Relase(ch)",
it will release that claim.

Caution:

Goroutines are created and submitted to the scheduler, but their first instructions are not run until the scheduler gives them a slot. 
This means that some goroutines might finish before others can run.
*/
package chancloser
