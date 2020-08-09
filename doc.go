/*
This package closes a channel when the last sender of that channel finishes work.
For this to work, a sender has to claim a channel with ChanClaim and using "defer ChanRelase" it will release that claim.
*/
package chancloser
