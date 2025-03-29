xivmount is a simple mount tracker for Final Fantasy XIV to check which Extreme/Savage mounts a character is still missing.

It has three main commands:

  go run . search \<World> \<name>
  
This command allows you to type in a world and character name to attempt to find their Lodestone ID. The Lodestone's character search functionality is notoriously poor, and since we're simply acting as a proxy for it, please expect lackluster results. I have set it to return results for "Level (decending)" by adding "order=7" to the url in search.go as it seems to be more accurate, though it can be changed to better serve your own needs. Not capitalizing first character of \<World> (ie Tonberry vs tonberry) will return results from every region. If you are getting strange results, this may be why. Additionally, you can use "_region_1" in \<World> to search by data center (1 is JP, 2 is NA, 3 is EU, 4 is OCE).

  go run . \<LodestoneID>
  
This command allows you to type in a character's Lodestone ID in order to see how many mounts they are missing compared to wishlist.go. By default, wishlist.go contains all the Extreme and Savage mounts, but any mount could be added to (or removed from) it if desired. Running the command creates a file saved at /character/<LodestoneID>.json which saves the character's Lodestone ID, Name, World, Avatar URL, and the remaining mounts. The Lodestone page updates sporatically (~12-24 hours), so to prevent unnecessary attempts and assist in rate-limiting, it will skip any Lodestone ID it's checked in the last 6 hours.

  go run . fc \<FC LodestoneID>
  
This command allows you to type in a Free Company's Lodestone ID in order to check every member in the FC against wishlist.go. This will create a file saved at /freecompany/<FC LodestoneID>.json that contains the Free Company Lodestone ID, Free Company Name, and a list of all member LodestoneIDs. It will also create a file in /character/<LodestoneID>.json for each member in the Free Company, also failing any members who were already checked within the last 6 hours.

Future Updates:
I am planning on adding the option to lookup Free Company Lodestone IDs similar to how you can look up Character Lodestone IDs. The Free Company LodestoneID command was more of a future-proofing / proof-of-concept feature, so I don't plan to add supporting commands at this time.
