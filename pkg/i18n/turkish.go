package i18n

func turkishSet() TranslationSet {
	return TranslationSet{
		PruningStatus:              "temizleniyor",
		RemovingStatus:             "kaldırılıyor",
		RestartingStatus:           "yeniden başlatılıyor",
		StoppingStatus:             "durduruluyor",
		RunningCustomCommandStatus: "özel komut çalıştır",

		NoViewMachingNewLineFocusedSwitchStatement: "NewLineFocused anahtar deyimi ile eşleşen görünüm yok",

		ErrorOccurred:                     "Bir hata oluştu! Lütfen https://github.com/jesseduffield/lazydocker/issues adresinden bir hataya ilişkin konu oluşturun",
		ConnectionFailed:                  "Docker bağlantısı başarısız oldu. Docker' ı yeniden başlatmanız gerekebilir",
		UnattachableContainerError:        "Konteyner attaching modunda çalışmayı desteklemiyor. Hizmeti '-it' opsiyonu ile çalıştırmanız veya docker-compose.yml dosyasında `stdin_open: true, tty: true` kullanmanız gerekir.",
		CannotAttachStoppedContainerError: "Durdurulan konteynera bağlanamazsınız, ilk önce başlatmanız gerekir (aslında başlatmayı r tuşu ile yapabilirsiniz) (evet, senin için bunu otomatik olarak yapabilirim fakat çok tembelim) (hata mesajı ile seninle birebir iletişim kurmam çok daha güzel)",
		CannotAccessDockerSocketError:     "Docker' a şu adresten erişilemiyor : unix:///var/run/docker.sock\n lazydocker' ı root(kök kullanıcı) olarak çalıştır veya şu adresteki adımları takip et : https://docs.docker.com/install/linux/linux-postinstall/",

		Donate:  "Bağış",
		Confirm: "Onayla",

		Return:             "dönüş",
		FocusMain:          "ana panele odaklan",
		Navigate:           "gezin",
		Execute:            "çalıştır",
		Close:              "kapat",
		Menu:               "menü",
		MenuTitle:          "Menü",
		Scroll:             "kaydır",
		OpenConfig:         "lazydocker ayarlarını aç",
		EditConfig:         "lazzydocker ayarlarını düzenle",
		Cancel:             "iptal",
		Remove:             "kaldır",
		ForceRemove:        "kaldırmaya zorla",
		RemoveWithVolumes:  "alanları ile birlikte kaldır",
		RemoveService:      "konteynerleri kaldır",
		Stop:               "durdur",
		Restart:            "yeniden başlat",
		Rebuild:            "yeniden yapılandır",
		Recreate:           "yeniden oluştur",
		PreviousContext:    "önceki sekme",
		NextContext:        "sonraki sekme",
		Attach:             "bağlan/iliştir",
		ViewLogs:           "kayıt defterini görüntüle",
		RemoveImage:        "imajı kaldır",
		RemoveVolume:       "alanı kaldır",
		RemoveNetwork:      "ağı kaldır",
		RemoveWithoutPrune: "etkisiz ebeveynleri silmeden kaldır",
		PruneContainers:    "çalışmayan konteynerleri temizle",
		PruneVolumes:       "kullanılmayan alanları temizle",
		PruneNetworks:      "kullanılmayan ağları temizle",
		PruneImages:        "kullanılmayan imajları temizle",
		ViewRestartOptions: "yeniden başlatma seçeneklerini görüntüle",
		RunCustomCommand:   "önceden tanımlanmış özel komutu çalıştır",

		GlobalTitle:               "Global",
		MainTitle:                 "Ana",
		ProjectTitle:              "Proje",
		ServicesTitle:             "Servisler",
		ContainersTitle:           "Konteynerler",
		StandaloneContainersTitle: "Bağımsız Konteynerler",
		ImagesTitle:               "Imajlar",
		VolumesTitle:              "Alanlar",
		NetworksTitle:             "Ağları",
		CustomCommandTitle:        "Özel Komut:",
		ErrorTitle:                "Hata",
		LogsTitle:                 "Kayitlar",
		ConfigTitle:               "Ayarlar",
		EnvTitle:                  "Env",
		DockerComposeConfigTitle:  "Docker-Compose Ayar",
		TopTitle:                  "Top",
		StatsTitle:                "Durumlar",
		CreditsTitle:              "Hakkinda",
		ContainerConfigTitle:      "Konteyner Ayar",
		ContainerEnvTitle:         "Konteyner Env",
		NothingToDisplay:          "Nothing to display",
		CannotDisplayEnvVariables: "Something went wrong while displaying environment variables",

		NoContainers: "Konteynerler yok",
		NoContainer:  "Konteyner yok",
		NoImages:     "Imajlar yok",
		NoVolumes:    "Alanlar yok",
		NoNetworks:   "Ağları yok",

		ConfirmQuit:                "Çıkmak istediğine emin misin?",
		MustForceToRemoveContainer: "Zorlamadan çalışan bir konteyneri kaldıramazsınız. Zorlamak ister misin?",
		NotEnoughSpace:             "Panelleri oluşturmak için yeterli alan yok",
		ConfirmPruneImages:         "Kullanılmayan tüm görüntüleri temizlemek istediğinize emin misiniz?",
		ConfirmPruneContainers:     "Durdurulan tüm konteynerları temizlemek istediğinizden emin misiniz?",
		ConfirmPruneVolumes:        "Kullanılmayan tüm alanları temizlemek istediğinizden emin misiniz?",
		ConfirmPruneNetworks:       "Kullanılmayan tüm ağları temizlemek istediğinizden emin misiniz?",
		StopService:                "Bu servisin konteynerlerini durdurmak istediğinize emin misiniz?",
		StopContainer:              "Bu konteyneri durdurmak istediğinize emin misiniz?",
		PressEnterToReturn:         "lazydocker' a geri dönmek için enter tuşuna basın ( Bu uyarı, `gui.return Immediately: true` ayarıyla devre dışı bırakılabilir)",
	}
}
