package main

// This package contains a "wishlist" of all the Extreme and Savage mounts players regularly farm for.
// It's used to compare against the scraped data in order to determine which mounts are still needed.
func LoadURLs() []string {
	return []string{
	// Horses
		// ifrit
		"https://lds-img.finalfantasyxiv.com/itemicon/98/98ca80ee41555010ea43b20e13cb09a88eb10447.png?n7.05",
		// garuda
		"https://lds-img.finalfantasyxiv.com/itemicon/25/25c30edf6a3823c579bd0834dcbb691fabcabd08.png?n7.05",
		// titan
		"https://lds-img.finalfantasyxiv.com/itemicon/b7/b76a674cb0a66f5eef23e1af99bc0ebe7243d344.png?n7.05",
		// shiva
		"https://lds-img.finalfantasyxiv.com/itemicon/29/29339eefcb3b0fa1ec0c9b42da2602f35cf1f277.png?n7.05",
		// ramuh
		"https://lds-img.finalfantasyxiv.com/itemicon/51/51b72bc49891c941fa12c5915c568cfdb38b6969.png?n7.05",
		// levi
		"https://lds-img.finalfantasyxiv.com/itemicon/c6/c68ae80fe089ceec3c19921cd94e96dbfa5e7772.png?n7.05",

	// Lanner
		// bismark
		"https://lds-img.finalfantasyxiv.com/itemicon/02/02bea2b33017af65733778de4598994bdbe1fa5f.png?n7.05",
		// ravana
		"https://lds-img.finalfantasyxiv.com/itemicon/d3/d33c9c54ea988b3d22b6336b22b5d5ab7f6bfe43.png?n7.05",
		// thordan
		"https://lds-img.finalfantasyxiv.com/itemicon/cd/cde21c7c5ff5979b2233ad38dc2f0d7302f61fbb.png?n7.05",
		// sephirot
		"https://lds-img.finalfantasyxiv.com/itemicon/c1/c1f202c000d50834e21e85d81a9ebd2a69467bb8.png?n7.05",
		// nidhogg
		"https://lds-img.finalfantasyxiv.com/itemicon/72/72833e97884788d44bdd233c007670751bcb624f.png?n7.05",
		// sophia
		"https://lds-img.finalfantasyxiv.com/itemicon/6a/6ae3a2779cc163397020d7a9218022074ac50dc1.png?n7.05",
		// zurvan
		"https://lds-img.finalfantasyxiv.com/itemicon/47/47d905eddc29499adfaff37042e4b4d6cdb6dcb4.png?n7.05",

	// Kamuy
		// lakshmi
		"https://lds-img.finalfantasyxiv.com/itemicon/09/095c4f4f58e0ae4424b79224815868560084dd38.png?n7.05",
		// susano
		"https://lds-img.finalfantasyxiv.com/itemicon/93/935e6c4141483878fd03d7b4d57047f00e9a8b90.png?n7.05",
		// shinryu
		"https://lds-img.finalfantasyxiv.com/itemicon/c9/c9a51f9e7ec8eb16fb26b549439cf55b7ee0cace.png?n7.05",
		// byakko
		"https://lds-img.finalfantasyxiv.com/itemicon/98/9894c0ce3662c696b271f82f3aa8ac4a2624d348.png?n7.05",
		// tsukuyomi
		"https://lds-img.finalfantasyxiv.com/itemicon/b7/b7d47b292f5a08608387574fce851caa36af2333.png?n7.05",
		// suzaku
		"https://lds-img.finalfantasyxiv.com/itemicon/32/32e0e5c87322cf3ce9cedade7987b5e06ed18938.png?n7.05",
		// seiryu
		"https://lds-img.finalfantasyxiv.com/itemicon/4d/4d83857e99d62d7b90753f412591fd4d9b634c4a.png?n7.05",

	// Gwiber
		// titania
		"https://lds-img.finalfantasyxiv.com/itemicon/93/93d92ad224adb1e8b339f7f050f5aa8353fdfcba.png?n7.05",
		// innocence
		"https://lds-img.finalfantasyxiv.com/itemicon/b1/b18730d0f3f0984758b6837d3029d245c7ed56b8.png?n7.05",
		// hades
		"https://lds-img.finalfantasyxiv.com/itemicon/f2/f2bbf659177044cc27bbeae1fd78db5c6885eed0.png?n7.05",
		// ruby
		"https://lds-img.finalfantasyxiv.com/itemicon/b6/b68b633a2e209bae6f33d5967dedf4e4f94a5350.png?n7.05",
		// sosex
		"https://lds-img.finalfantasyxiv.com/itemicon/4f/4f1db7e9bcb5d78749b468e092bae850d1ff8e76.png?n7.05",
		// emerald
		"https://lds-img.finalfantasyxiv.com/itemicon/b9/b931b1725a74e5cc95d5c8b1fb1d56b05ca26b55.png?n7.05",
		// diamond
		"https://lds-img.finalfantasyxiv.com/itemicon/1b/1b6d47ce44210e7c8635109f86d3b7c3073b2b56.png?n7.05",

	// Lynx
		// zodex
		"https://lds-img.finalfantasyxiv.com/itemicon/bb/bb2d2c0b85af0d2bf6dce40753ef6ce0f3af8bcc.png?n7.05",
		// hydex
		"https://lds-img.finalfantasyxiv.com/itemicon/03/0360370898e017488972f35aad2753faec2854cd.png?n7.05",
		// endex
		"https://lds-img.finalfantasyxiv.com/itemicon/e9/e91053a91bb158ac902d4929fef7a8dc2f229960.png?n7.05",
		// barbariccia
		"https://lds-img.finalfantasyxiv.com/itemicon/53/5351b041a664233d95ba2223945a2815422dc5fe.png?n7.05",
		// rubicante
		"https://lds-img.finalfantasyxiv.com/itemicon/13/13d1f541dbb76299473cc8bc013b4253a50919fd.png?n7.05",
		// golbez
		"https://lds-img.finalfantasyxiv.com/itemicon/ba/ba2404c7381d10bec9f24059976aff8258bcb9c5.png?n7.05",
		// zeromus
		"https://lds-img.finalfantasyxiv.com/itemicon/ec/ec4bed15f67129295a4211027292dbc631e17783.png?n7.05",

	// Wings
		// 1
		"https://lds-img.finalfantasyxiv.com/itemicon/86/86ed3b6e18652430d112de4ee4fc504ea6ec088d.png?n7.05",
		// 2
		"https://lds-img.finalfantasyxiv.com/itemicon/4d/4d3a9eda7aa3c3817107354d2e74fc07a93d20d5.png?n7.05",

	// Savage
		// A4S
		"https://lds-img.finalfantasyxiv.com/itemicon/09/09e6146102d51f40e5811b5c00f13830e668e709.png?n7.05",
		// A12S
		"https://lds-img.finalfantasyxiv.com/itemicon/54/541f463a23864991a97a5a9fc4c5674bc4f66525.png?n7.05",
		
		// O4S
		"https://lds-img.finalfantasyxiv.com/itemicon/82/823e778dd0a6e5c2c923149b5f1b62793704b8e2.png?n7.05",
		// O8S
		"https://lds-img.finalfantasyxiv.com/itemicon/e5/e576570acd6ba85b627069e629ded31fec05aef6.png?n7.05",
		// O12S
		"https://lds-img.finalfantasyxiv.com/itemicon/4e/4ec48d8928667db38d951562230ae3dd97ac152b.png?n7.05",
		
		// E4S
		"https://lds-img.finalfantasyxiv.com/itemicon/36/3622191d64675b460fba406b6814348c945e5cfb.png?n7.05",
		// E8S
		"https://lds-img.finalfantasyxiv.com/itemicon/6c/6c7a179ecc39a2c8750bd1a751152ac9853d345d.png?n7.05",
		// E12S
		"https://lds-img.finalfantasyxiv.com/itemicon/bf/bf69c30fa1c4da30c98ff0c64baef072abede17e.png?n7.05",

		// P4S
		"https://lds-img.finalfantasyxiv.com/itemicon/cc/cc4c31253e3089ae74bc5416726a4d91751f2a30.png?n7.05",
		// P8S
		"https://lds-img.finalfantasyxiv.com/itemicon/76/763a6d3e9bfb94764eb2f65ffb53108d6041f919.png?n7.05",
		// P12S
		"https://lds-img.finalfantasyxiv.com/itemicon/32/32e8c06c3e74bbb0556bed939312744c19789d3f.png?n7.05",

		// M4S

		// M8S

		// M12S
		
	}
}
