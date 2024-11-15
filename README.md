xivmount simple mount tracker for Final Fantasy XIV to check which Extreme/Savage mounts a character is still missing.

It has three main commands:
  go run . search <World> <Name>
This command allows you to type in a world and character name to attempt to find their Lodestone ID.
The Lodestone's character search functionality is notoriously poor, and since we're simply coopting it, please expect lackluster results.
Returns results as "Level (decending)" by adding "order=7" to the url in search.go for accuracy. I am planning on making this variable in future.
The world won't work unless the first character is capitalized (ie Tonberry > tonberry). This is how their url works and not something I feel like addressing.
You can use "_region_1" in <World> to search by data center. 1 is JP, 2 is NA, 3 is EU, 4 is OCE.

  go run . <LodestoneID>
This command allows you to type in a character's Lodestone ID in order to see how many mounts they are missing compared to wishlist.go
By default, wishlist.go contains all the Extreme and Savage mounts, but any mount could be added to it if desired.
Running the command creates a file saved at /character/<LodestoneID>.json which saves the character's Lodestone ID, Name, World, Avatar URL, and the remaining mounts.
The Lodestone page updates sporatically (~12-24 hours), so to prevent unnecessary attempts, it will fail to check any Lodestone ID it's checked in the last 6 hours.

  go run . fc <FC LodestoneID>
This command allows you to type in a Free Company's Lodestone ID in order to check every member in the FC against wishlist.go
This will create a file saved at /freecompany/<FC LodestoneID>.json that contains the Free Company Lodestone ID, Free Company Name, and a list of all member LodestoneIDs.
It will also create a file in /character/ for each member in the Free Company, failing any members who were checked within the last 6 hours.

Future Updates:
I am planning on adding the option to lookup FC Lodestone IDs similar to how you can look up Character Lodestone IDs.
