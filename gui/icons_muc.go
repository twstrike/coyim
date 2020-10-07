package gui

import (
	"fmt"

	"github.com/coyim/gotk3adapter/gdki"
)

var mucIcon = map[string]*icon{
	"room": &icon{
		size:   "32x32",
		width:  32,
		height: 32,

		encoded: "" +
			"89504e470d0a1a0a0000000d4948445200000020000000200806000000737a7af400000009704859" +
			"7300000b1300000b1301009a9c18000000017352474200aece1ce90000000467414d410000b18f0b" +
			"fc61050000050e494441547801cd564d6c1b55109e79bba9dda2a4768228a84831554a25ca4f0442" +
			"88036de09403524d102807daa6856b85c58d53e35c7a24074851a550a78250e5804a38c1a5e55471" +
			"69c30968054e0a5c802406159c38d937ccccdb676fdca8d9434479f2da6f77df9bef9b996fe619e0" +
			"1e0f94afe5d1522eb3333c4d404544e8256abddb6407f929af43bf88106a68e15283a272feddf179" +
			"484b60f94ca9b003f032220ab0da13047440a830eeb1c32268326882c773dd43b0b00af6c5b4244c" +
			"c6e211b05020cb3ef017f0c71083467ccf73bd641ec5736149f13cbef773cbebaca54287c512a41c" +
			"c612159d119e3912c203c9dad8303bd69583e0d1c7419f756420e83b28eb819ae0d611531656b274" +
			"243501d92886159cedb00d6a79669558c7b38720e48b3da7b0b70f32afbde900ad0252d386444af6" +
			"45362d3e847eb1e48f07ba848bba10c2e70e0366776270e049b0b545e87861104def7ea0da12741c" +
			"1a04a82de1dadc377718c58450b72640cedb76e1b36b101c78c23dcd7503aefc03017b8f0fee0558" +
			"a9eb9cf6d401ae5d25888ba1a5584c8b2f11607014e1692929b45335c2cae438848f3d0541613fd4" +
			"3f384398efc15def8c417dfa1cd8ea0def2efaba9008ea2f626a06a1880e357571b5f1c4a82b7acf" +
			"21ef83a87a53c516ecd9ab9bec4f379a7152cd78221a49de841b35f0dd70b1c06bab8947958c0946" +
			"f74d7fb610b270780339a75d2ad0e47b68c74b2fab17c13ece79bd0e99578fa179e861200e7f66e8" +
			"28444c62edda558556cd50a22b6ca60022272cb1893812813d7c7da4f8745c05aee6c9b6d57736cb" +
			"f9ef01bbbce8549fef065a5a6cbde30a40ad1a00b7dfea33b5b3c910a60c0e86af00b0d0dd306f87" +
			"5a6ae853805a6af68fdf6175a602e1c17ed6403fac4c4d7035ec82fb9e791e56bf384bd18f3f80f7" +
			"c6e59d126d13d120ddc1008dc1209687d18b1f010e3001a849b771f9776b950cdf444ca4313b237d" +
			"018568e3ab59886e7eefc0c553512ab98d4e02b1080dfcd94e80c1636054a840e09d08e95bde5274" +
			"e0e47f3428f6d75bd8f8e596dbf1f76d6a7c398b5e2bcd43a355012d0d44309704cff239638d743d" +
			"894e3302ee5ab7ebe3eccd1cb4e5dfebc2f77ac54cf47d7d1f69d7a3e43ad6c4fc9a8dca1bdc3738" +
			"e28103b6e27e5507ad8e71fbd4c9122bea382729d7943169800b120f9f63ffc3ebe6212170f6a6c6" +
			"59fd7c1dec787ebc52f3efaa5c821818396d0b2a405e687c3a00af6cd930fe3a35525512719f88db" +
			"43b9f3fdca28a4180b4787de33604a9a77511e681454331c812be19616229becb1129172e7c45439" +
			"0df8cf6f0c9de65d25cd35c4b907f1de682ac4e0d604b420c8650b61ac73e2c296e0d5623117749a" +
			"18dc38501f7ae3ee115d1da78840e468181ceb3cfbf168f2d57506ca66d7fb433613b24c2203850e" +
			"c07e96fa0813ce055eeda0c24325a01168492a450458e4ecf9ee0f3f19dd003e3c58d88974d940d8" +
			"2bee90e6d495af969b07d572c356138abd77a6d310005bde7dee6225f9440e17c661702c88a9c0d7" +
			"b4cbb3af759408b8f76dc0f22746e758497f70b7831b2e2b3df970a3ba21ae7563148dd7e8e9eac1" +
			"d1b77dc04b5d93d3afa488c0c6c1c6fad9d379b6336f6240c9686060803d25efa9af75fe7c8d71c3" +
			"913ecbd8d223047c4a09c1368ddf8ebd7e9cad9d97f81bf55ca3b0d035f9e92377db67609bc60317" +
			"66a6b8c24e04fec83526ee1e771fdb4640c6fde767a638fa27e313aff95ff13f232023ffd1c54a44" +
			"748252a677db092449c0bd1ecb6f0d0fc0ff7dfc0b0cc18fd12a92766b0000000049454e44ae4260" +
			"82",
	},
	"occupant": &icon{
		size:   "24x24",
		width:  24,
		height: 24,
		encoded: "" +
			"89504e470d0a1a0a0000000d49484452000000180000001b080600000066e34f560000000970485" +
			"97300000b1300000b1301009a9c18000000017352474200aece1ce90000000467414d410000b18f" +
			"0bfc6105000005014944415478019d565b6c145518feff33b3db76d9b6db962d9796b214b584c60" +
			"09a189fd0468d68d43631f55143e293518a6fa0c2ce86883e682caf10138d26461f146d045e085b" +
			"5f8c26255e5a4485766d49b96c69b7f6bedb99dfff3f6766b7b594269ccdd9993997effbfedb990" +
			"1b8b7867e5fb3d9771a24fe2121ca35954aa1c3634ed2cc39fe1a272563495ec14c8866db2a4afe" +
			"8f8e9002ecdcd989ad9daf61cfc92fb0b966334e476b71ae621c2193818a0d95149daea3c1e64d5" +
			"439759de2d9016abdd44ac9649234a0685b85003bbfea54d9782bc6b3a0661250bb2e54f63a11ec" +
			"45544dccbe55d4b2861c78eeefbcbb0794fa6ee0cfb9a1fa3878e95ef020e9908f4ccb08e429458" +
			"eeae91bb59adcc4363b84a758c95e22f2cd42ed3a323e61a52216c110c2e7167ac76e4ee4078b24" +
			"3e815594ee382a0da01a2be20794ad3e63a81601e755441ea1b937c09a52c6029504bb78eee548b" +
			"9bd50bfd1fd39f3e9e340e90b94625f2b239f7060e700aeaf0e1df5c8fbd073dd9806f478bf2730" +
			"a29c9beba1e77ac057d0a6781e1277de23f3d5dc3fa8ad0c1d85a4b14da055100a77fb9e3778cf1" +
			"15f32efe52ea072e59b96750df8d486ddd0def028df6f46191301dc098c10ad83818fbc78b1ac2b" +
			"c829cdf2fc4fc7b695d9561f2fa821a3060c1141c42a83030fbc003baab62ccb868be357e0d4e03" +
			"99877f3121393aa7e4451618e5cf5f0d78f1c1ed416841193ac28266a8ddb918c7a82579b9f5e01" +
			"2eeda1dafba0ebfe7663a5f16049980731ce838f3559c78f4e42d9a1418983ce123157f2987ddb5" +
			"2b9050eb7be04776bef0d7c0997ff1d61d52a28b820b380428b35cab242224367092758d135b260" +
			"475523acd59a22f5c6dd1e4131eb88b4b754de3a68bb9ed7619652f06fa2c51b660af36b1254586" +
			"12d8cfc7c4411ab44af4ee2c76cfedb654e1f534c92e79e67c886676ead49707972a4f420292dd2" +
			"3d290d5d8909e6a59a203826dfcd55fec5b77f4c0eaf0ade77fb0acf8fe84a235f9481d28e92d18" +
			"4927c2f79a6b410c814d84797be61a0bf57800bf0c9bfbe97ba360143307b03b1609ee5b8cef0f3" +
			"56939f828a26e0dc94b2a8ae228a3f8cfd0643b3a3d0188973de17606c210757a746211a2e87b9b" +
			"93c723d6bdff38108c15e9f74c266960c9b98d089e5bf03646134540e0fd635a1ad2c90a320bb300" +
			"9b718589aa461ac3c0ab5916a68c8cf42ffd83f5820b7689d3ebcb450f85531701a7ccb82f3b25cd" +
			"9b0a7be19d6b1c290659b6edb10b642fa3e1c8c31f9fa8a2a5ebb5dc776c929a0bdc0eeff56a9457" +
			"5422a18f4712291414cc43612bf07c0624b2c2e209b09c34a402d03cef762992d73bc4648ea23312" +
			"8258ba9a5906b9d56a7db9c1c3fa497320fe56ec0c262019410f82416835b2ae84a5f655ef1fdb5e" +
			"9db707d7adc94987f64b0459f30764697874b8bfb193c174cce7281f50ef7c3b5a9319f0475d78a9" +
			"54fc8cf129bfeec30fc72e3aa79538803f4110bb9825b48012c79a33dd7fb4e17cf759b273476324" +
			"8c40e431d077353b446fb5cda6c6101b2735374737a02f36c296071875f67f0e6d9b677bb9711487" +
			"bf6fc5b0e8f244bcbc18fbc7f8a7116ea6a85a55b0996ade70f91334f1c7702cc155f15cf9c7ffb2" +
			"063267979cc4f2b2c119903d32f74b31dc97fb9e3244b77ce3c79fcc452bc3b7e3cedbb7028012e2" +
			"419ed9522c1b25dfe10150fc83458b4ff5cdbfb99154be12e6ddfd94309b4b18343d7ce0a77f350c" +
			"c87ccf0d60c13f4cedbf3dde9b6eedc6a18ff01decce133ca3521200000000049454e44ae426082",
	},
}

func getMUCIcon(name string) Icon {
	ico, ok := mucIcon[name]
	if !ok {
		panic(fmt.Sprintf("getMUCIcon(): icon %s doesn't exists", name))
	}
	return ico
}

func getMUCIconPixbuf(name string) gdki.Pixbuf {
	ico := getMUCIcon(name)
	return ico.GetPixbuf()
}
