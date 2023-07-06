package macinfo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type SPAirPortDataType struct {
	SPAirPortSoftwareInformation []SPAirPortSoftwareInformation `json:"SPAirPortDataType"`
}

type SPAirPortSoftwareInformation struct {
	SPAirPortCoreWLANVersion    string `json:"spairport_corewlan_version"`
	SPAirPortCoreWLANKitVersion string `json:"spairport_corewlankit_version"`
	SPAirPortDiagnosticsVersion string `json:"spairport_diagnostics_version"`
	SPAirPortExtraVersion       string `json:"spairport_extra_version"`
	SPAirPortFamilyVersion      string `json:"spairport_family_version"`
	SPAirPortProfilerVersion    string `json:"spairport_profiler_version"`
	SPAirPortUtilityVersion     string `json:"spairport_utility_version"`
}

func getSPAirportDataType() (result SPAirPortDataType, fbErr error) {
	defer func() {
		if er := recover(); er != nil {
			fbErr = errors.New("Critical error with getting SPAirPortDataType")
		}
	}()
	var res SPAirPortDataType
	command := exec.Command("system_profiler", "SPAirPortDataType", "-json")
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return res, err
	} else {
		if err = json.Unmarshal(out.Bytes(), &result); err != nil {
			return res, err
		} else {
			return result, nil
		}
	}
}

// geting
type SPApplicationsDataType struct {
	Applications []SPApplication `json:"SPApplicationsDataType"`
}

type SPApplication struct {
	Name         string    `json:"_name"`
	ArchKind     string    `json:"arch_kind"`
	LastModified time.Time `json:"lastModified"`
	ObtainedFrom string    `json:"obtained_from"`
	Path         string    `json:"path"`
	SignedBy     []string  `json:"signed_by"`
	Version      string    `json:"version"`
}

func getSPApplicationsDataType() (result SPApplicationsDataType, fbErr error) {
	defer func() {
		if er := recover(); er != nil {
			fbErr = errors.New("Critical error with getting SPApplications")
		}
	}()
	var res SPApplicationsDataType
	command := exec.Command("system_profiler", "SPApplicationsDataType", "-json")
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return res, err
	} else {
		if err = json.Unmarshal(out.Bytes(), &result); err != nil {
			return res, err
		} else {
			return result, nil
		}
	}
}

// getting
type SPAudioDataType struct {
	AudioItems []SPAudioItem `json:"SPAudioDataType"`
}

func getSPAudioDataType() (result SPAudioDataType, fbErr error) {
	defer func() {
		if er := recover(); er != nil {
			fbErr = errors.New("Critical error with getting SPAudioDataType")
		}
	}()
	var res SPAudioDataType
	command := exec.Command("system_profiler", "SPAudioDataType", "-json")
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return res, err
	} else {
		if err = json.Unmarshal(out.Bytes(), &result); err != nil {
			return res, err
		} else {
			return result, nil
		}
	}
}

type SPAudioItem struct {
	Name                             string `json:"_name"`
	CoreAudioDeviceInput             int    `json:"coreaudio_device_input"`
	CoreAudioDeviceManufacturer      string `json:"coreaudio_device_manufacturer"`
	CoreAudioDeviceSampleRate        int    `json:"coreaudio_device_srate"`
	CoreAudioDeviceTransport         string `json:"coreaudio_device_transport"`
	CoreAudioInputSource             string `json:"coreaudio_input_source"`
	CoreAudioDefaultAudioInputDevice string `json:"coreaudio_default_audio_input_device"`
}

//gettinh
//Блютуз - на потом
/*
type SPBluetoothDataType struct {
	Items []SPBluetoothItem `json:"SPBluetoothDataType"`
}

type SPBluetoothItem struct {
	AppleBluetoothVersion    string               `json:"apple_bluetooth_version"`
	IncomingSerialPortsTitle []IncomingSerialPort `json:"incoming_serial_ports_title"`
	LocalDeviceTitle         LocalDevice          `json:"local_device_title"`
	ServicesTitle            []Service            `json:"services_title"`
}

type IncomingSerialPort struct {
	BluetoothIncomingPort BluetoothIncomingPort `json:"Bluetooth-Incoming-Port"`
}

type BluetoothIncomingPort struct {
	DeviceAuthentication string `json:"device_authentication"`
	DeviceChannel        string `json:"device_channel"`
}

type LocalDevice struct {
	GeneralAddress                string `json:"general_address"`
	GeneralAutoSeekKeyboard       string `json:"general_autoseek_keyboard"`
	GeneralAutoSeekPointing       string `json:"general_autoseek_pointing"`
	GeneralDeviceClassComposite   string `json:"general_device_class_composite"`
	GeneralDeviceClassMajor       string `json:"general_device_class_major"`
	GeneralDeviceClassMinor       string `json:"general_device_class_minor"`
	GeneralFwVersion              string `json:"general_fw_version"`
	GeneralHardwareTransport      string `json:"general_hardware_transport"`
	GeneralHciRevision            string `json:"general_hci_revision"`
	GeneralHciVersion             string `json:"general_hci_version"`
	GeneralLmpSubversion          string `json:"general_lmp_subversion"`
	GeneralLmpVersion             string `json:"general_lmp_version"`
	GeneralMfg                    string `json:"general_mfg"`
	GeneralName                   string `json:"general_name"`
	GeneralPower                  string `json:"general_power"`
	GeneralProductID              string `json:"general_productID"`
	GeneralRemoteWake             string `json:"general_remoteWake"`
	GeneralServiceClass           string `json:"general_service_class"`
	GeneralSupportsHandoff        string `json:"general_supports_handoff"`
	GeneralSupportsInstantHotspot string `json:"general_supports_instantHotspot"`
	GeneralSupportsLowEnergy      string `json:"general_supports_lowEnergy"`
	GeneralTypeCompleteString     string `json:"general_type_complete_string"`
	GeneralTypeMajorString        string `json:"general_type_major_string"`
	GeneralVendorID               string `json:"general_vendorID"`
}

type Service struct {
	FileBrowsingTitle    FileBrowsing    `json:"file_browsing_title,omitempty"`
	ObjectPushTitle      ObjectPush      `json:"object_push_title,omitempty"`
	InternetSharingTitle InternetSharing `json:"internet_sharing_title,omitempty"`
}

type FileBrowsing struct {
	ServiceFTPRootFolder string `json:"service_FTPRootFolder"`
	ServiceOBEXReceive   string `json:"service_OBEXReceive"`
	ServiceState         string `json:"service_state"`
}

type ObjectPush struct {
	ServiceOBEXFolder      string `json:"service_OBEXFolder"`
	ServiceOBEXHandleOther string `json:"service_OBEXHandle_Other"`
	ServiceOBEXReceive     string `json:"service_OBEXReceive"`
	ServiceState           string `json:"service_state"`
}

type InternetSharing struct {
	ServiceState string `json:"service_state"`
}

/*
type SPBluetoothDataType struct {
	Items []SPBluetoothItem `json:"SPBluetoothDataType"`
}

type SPBluetoothItem struct {
	AppleBluetoothVersion    string           `json:"apple_bluetooth_version"`
	IncomingSerialPortsTitle []string         `json:"incoming_serial_ports_title"`
	LocalDeviceTitle         LocalDeviceTitle `json:"local_device_title"`
	ServicesTitle            []ServicesTitle  `json:"services_title"`
}
*/
/*
type LocalDeviceTitle struct {
	GeneralAddress            string `json:"general_address"`
	GeneralName               string `json:"general_name"`
	GeneralTypeCompleteString string `json:"general_type_complete_string"`
}

type ServicesTitle struct {
	FileBrowsingTitle    FileBrowsingTitle    `json:"file_browsing_title"`
	ObjectPushTitle      ObjectPushTitle      `json:"object_push_title"`
	InternetSharingTitle InternetSharingTitle `json:"internet_sharing_title"`
}

type FileBrowsingTitle struct {
	ServiceFTPRootFolder string `json:"service_FTPRootFolder"`
	ServiceOBEXReceive   string `json:"service_OBEXReceive"`
	ServiceState         string `json:"service_state"`
}

type ObjectPushTitle struct {
	ServiceOBEXFolder      string `json:"service_OBEXFolder"`
	ServiceOBEXHandleOther string `json:"service_OBEXHandle_Other"`
	ServiceOBEXReceive     string `json:"service_OBEXReceive"`
	ServiceState           string `json:"service_state"`
}

type InternetSharingTitle struct {
	ServiceState string `json:"service_state"`
}
*/
type SPDisplaysDataType struct {
	SPDisplays []SPDisplay `json:"SPDisplaysDataType"`
}

type SPDisplay struct {
	Name          string `json:"_name"`
	DeviceID      string `json:"spdisplays_device-id"`
	Ndrvs         []Ndrv `json:"spdisplays_ndrvs"`
	RevisionID    string `json:"spdisplays_revision-id"`
	VendorID      string `json:"spdisplays_vendor-id"`
	VRAM          string `json:"spdisplays_vram"`
	PciDeviceType string `json:"sppci_device_type"`
}

type Ndrv struct {
	Name              string `json:"_name"`
	ProductID         string `json:"_spdisplays_display-product-id"`
	SerialNumber2     string `json:"_spdisplays_display-serial-number2"`
	VendorID          string `json:"_spdisplays_display-vendor-id"`
	DisplayID         string `json:"_spdisplays_displayID"`
	DisplayPath       string `json:"_spdisplays_displayPath"`
	DisplayRegID      string `json:"_spdisplays_displayRegID"`
	Pixels            string `json:"_spdisplays_pixels"`
	Resolution        string `json:"_spdisplays_resolution"`
	AmbientBrightness string `json:"spdisplays_ambient_brightness"`
	Depth             string `json:"spdisplays_depth"`
	Main              string `json:"spdisplays_main"`
	Mirror            string `json:"spdisplays_mirror"`
	Online            string `json:"spdisplays_online"`
	PixelResolution   string `json:"spdisplays_pixelresolution"`
	DisplayResolution string `json:"spdisplays_resolution"`
}

// getting
type SPEthernetDataType struct {
	Ethernet []SPEthernet `json:"SPEthernetDataType"`
}

type SPEthernet struct {
	Name             string `json:"_name"`
	BSDName          string `json:"spethernet_BSD_Name"`
	BUNDLEIdentifier string `json:"spethernet_BUNDLE_IDentifier"`
	Bus              string `json:"spethernet_bus"`
	DeviceType       string `json:"spethernet_device_type"`
	DeviceID         string `json:"spethernet_device-id"`
	KextPath         string `json:"spethernet_kext_path"`
	LinkWidth        string `json:"spethernet_link-width"`
	SpEtherName      string `json:"spethernet_name"`
	RevisionID       string `json:"spethernet_revision-id"`
	VendorID         string `json:"spethernet_vendor-id"`
	Version          string `json:"spethernet_version"`
}

// getting
type SPHardwareDataType struct {
	Hardware []SPHardware `json:"SPHardwareDataType"`
}

func getSPHardwareDataType() (result SPHardwareDataType, fbErr error) {
	defer func() {
		if er := recover(); er != nil {
			fbErr = errors.New("Critical error with getting SPHardwareDataType")
		}
	}()
	var res SPHardwareDataType
	command := exec.Command("system_profiler", "SPHardwareDataType", "-json")
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return res, err
	} else {
		if err = json.Unmarshal(out.Bytes(), &result); err != nil {
			return res, err
		} else {
			return result, nil
		}
	}
}

type SPHardware struct {
	Name                  string `json:"_name"`
	AppleROMInfo          string `json:"apple_rom_info"`
	BootROMVersion        string `json:"boot_rom_version"`
	CurrentProcessorSpeed string `json:"current_processor_speed"`
	L2Cache               string `json:"l2_cache"`
	MachineModel          string `json:"machine_model"`
	MachineName           string `json:"machine_name"`
	//NumberProcessors      int    `json:"number_processors"`
	NumberProcessors interface{} `json:"number_processors"`
	Packages         int         `json:"packages"`
	PhysicalMemory   string      `json:"physical_memory"`
	PlatformUUID     string      `json:"platform_UUID"`
	SerialNumber     string      `json:"serial_number"`
	SMCVersionSystem string      `json:"SMC_version_system"`
}

// getting
type SPMemoryDataType struct {
	Memory []SPMemory `json:"SPMemoryDataType"`
}

type SPMemory struct {
	Items               []SPMemoryItem `json:"_items"`
	Name                string         `json:"_name"`
	GlobalECCState      string         `json:"global_ecc_state"`
	IsMemoryUpgradeable string         `json:"is_memory_upgradeable"`
	DimmManufacturer    string         `json:"dimm_manufacturer"`
	DimmType            string         `json:"dimm_type"`
}

type SPMemoryItem struct {
	Name             string `json:"_name"`
	DIMMManufacturer string `json:"dimm_manufacturer"`
	DIMMPartNumber   string `json:"dimm_part_number"`
	DIMMSerialNumber string `json:"dimm_serial_number"`
	DIMMSize         string `json:"dimm_size"`
	DIMMSpeed        string `json:"dimm_speed"`
	DIMMStatus       string `json:"dimm_status"`
	DIMMType         string `json:"dimm_type"`
}

// getting
type SPNetworkDataType struct {
	Network []SPNetwork `json:"SPNetworkDataType"`
}

type SPNetwork struct {
	Name                  string            `json:"_name"`
	DHCP                  SPNetworkDHCP     `json:"dhcp"`
	DNS                   SPNetworkDNS      `json:"DNS"`
	Ethernet              SPNetworkEthernet `json:"Ethernet"`
	Hardware              string            `json:"hardware"`
	Interface             string            `json:"interface"`
	IPAddress             []string          `json:"ip_address"`
	IPv4                  SPNetworkIPv4     `json:"IPv4"`
	IPv6                  SPNetworkIPv6     `json:"IPv6"`
	Proxies               SPNetworkProxies  `json:"Proxies"`
	SPNetworkServiceOrder int               `json:"spnetwork_service_order"`
	Type                  string            `json:"type"`
}

type SPNetworkDHCP struct {
	DomainName        string `json:"dhcp_domain_name"`
	DomainNameServers string `json:"dhcp_domain_name_servers"`
	LeaseDuration     int    `json:"dhcp_lease_duration"`
	MessageType       string `json:"dhcp_message_type"`
	Routers           string `json:"dhcp_routers"`
	ServerIdentifier  string `json:"dhcp_server_identifier"`
	SubnetMask        string `json:"dhcp_subnet_mask"`
}

type SPNetworkDNS struct {
	DomainName      string   `json:"DomainName"`
	ServerAddresses []string `json:"ServerAddresses"`
}

type SPNetworkEthernet struct {
	MACAddress   string   `json:"MAC Address"`
	MediaOptions []string `json:"MediaOptions"`
	MediaSubType string   `json:"MediaSubType"`
}

type SPNetworkIPv4 struct {
	AdditionalRoutes           []SPNetworkAdditionalRoute `json:"AdditionalRoutes"`
	Addresses                  []string                   `json:"Addresses"`
	ARPResolvedHardwareAddress string                     `json:"ARPResolvedHardwareAddress"`
	ARPResolvedIPAddress       string                     `json:"ARPResolvedIPAddress"`
	ConfigMethod               string                     `json:"ConfigMethod"`
	ConfirmedInterfaceName     string                     `json:"ConfirmedInterfaceName"`
	InterfaceName              string                     `json:"InterfaceName"`
	NetworkSignature           string                     `json:"NetworkSignature"`
	Router                     string                     `json:"Router"`
	SubnetMasks                []string                   `json:"SubnetMasks"`
}

type SPNetworkAdditionalRoute struct {
	DestinationAddress string `json:"DestinationAddress"`
	SubnetMask         string `json:"SubnetMask"`
}

type SPNetworkIPv6 struct {
	ConfigMethod string `json:"ConfigMethod"`
}

type SPNetworkProxies struct {
	ExceptionsList []string `json:"ExceptionsList"`
	FTPPassive     string   `json:"FTPPassive"`
}

// getting
type SPPrintersDataType struct {
	Printers []SPPrinter `json:"SPPrintersDataType"`
}

type SPPrinter struct {
	Name                   string            `json:"_name"`
	AirPrintVersion        string            `json:"airprintversion"`
	CreationDate           string            `json:"creationDate"`
	CupsFilters            []SPPrinterFilter `json:"cups filters"`
	CupsVersion            string            `json:"cupsversion"`
	Default                string            `json:"default"`
	DriverVersion          string            `json:"driverversion"`
	FaxSupport             string            `json:"Fax Support"`
	PPD                    string            `json:"ppd"`
	PPDFileVersion         string            `json:"ppdfileversion"`
	PrinterCommands        string            `json:"printercommands"`
	PrinterFirmwareVersion string            `json:"printerfirmwareversion"`
	PrinterPDEs            []string          `json:"printerpdes"`
	PrinterSharing         string            `json:"printersharing"`
	PrintServer            string            `json:"printserver"`
	PSVersion              string            `json:"psversion"`
	Scanner                string            `json:"scanner"`
	Shared                 string            `json:"shared"`
	Status                 string            `json:"status"`
	URFVersion             string            `json:"urfversion"`
	URI                    string            `json:"uri"`
}

type SPPrinterFilter struct {
	Name  string `json:"-"`
	Value string `json:"-"`
}

// getting
type SPSerialATADataType struct {
	Items []SPSerialATAItem `json:"SPSerialATADataType"`
}

type SPSerialATAItem struct {
	Name             string              `json:"_name"`
	Items            []SPSerialATAItem   `json:"_items"`
	BSDName          string              `json:"bsd_name"`
	DetachableDrive  string              `json:"detachable_drive"`
	DeviceModel      string              `json:"device_model"`
	DeviceRevision   string              `json:"device_revision"`
	DeviceSerial     string              `json:"device_serial"`
	PartitionMapType string              `json:"partition_map_type"`
	RemovableMedia   string              `json:"removable_media"`
	Size             string              `json:"size"`
	SizeInBytes      int64               `json:"size_in_bytes"`
	SPSataMediumType string              `json:"spsata_medium_type"`
	SPSataNCQ        string              `json:"spsata_ncq"`
	SPSataNCQDepth   string              `json:"spsata_ncq_depth"`
	Volumes          []SPSerialATAVolume `json:"volumes"`
}

type SPSerialATAVolume struct {
	Name        string `json:"_name"`
	BSDName     string `json:"bsd_name"`
	FileSystem  string `json:"file_system"`
	IOContent   string `json:"iocontent"`
	Size        string `json:"size"`
	SizeInBytes int64  `json:"size_in_bytes"`
	VolumeUUID  string `json:"volume_uuid,omitempty"`
}

type SPSoftwareDataType struct {
	Items []SPSoftware `json:"SPSoftwareDataType"`
}

type SPSoftware struct {
	Name            string `json:"_name"`
	BootMode        string `json:"boot_mode"`
	BootVolume      string `json:"boot_volume"`
	KernelVersion   string `json:"kernel_version"`
	LocalHostName   string `json:"local_host_name"`
	OSVersion       string `json:"os_version"`
	SecureVM        string `json:"secure_vm"`
	SystemIntegrity string `json:"system_integrity"`
	Uptime          string `json:"uptime"`
	UserName        string `json:"user_name"`
}

type macInfo struct {
	SPAirPortSoftwareInformation []SPAirPortSoftwareInformation `json:"SPAirPortDataType"`
	Applications                 []SPApplication                `json:"SPApplicationsDataType"`
	AudioItems                   []SPAudioItem                  `json:"SPAudioDataType"`
	SPDisplays                   []SPDisplay                    `json:"SPDisplaysDataType"`
	Ethernet                     []SPEthernet                   `json:"SPEthernetDataType"`
	Hardware                     []SPHardware                   `json:"SPHardwareDataType"`
	Memory                       []SPMemory                     `json:"SPMemoryDataType"`
	Network                      []SPNetwork                    `json:"SPNetworkDataType"`
	Printers                     []SPPrinter                    `json:"SPPrintersDataType"`
	SystemItems                  []SPSoftware                   `json:"SPSoftwareDataType"`
	SataItems                    []SPSerialATAItem              `json:"SPSerialATADataType"`
}

type Volume struct {
	Name             string `json:"_name"`
	BsdName          string `json:"bsd_name"`
	FileSystem       string `json:"file_system,omitempty"`
	IOContent        string `json:"iocontent,omitempty"`
	Size             string `json:"size"`
	SizeInBytes      int64  `json:"size_in_bytes"`
	VolumeUUID       string `json:"volume_uuid,omitempty"`
	FreeSpace        string `json:"free_space,omitempty"`
	FreeSpaceInBytes int64  `json:"free_space_in_bytes,omitempty"`
	MountPoint       string `json:"mount_point,omitempty"`
	Writable         string `json:"writable,omitempty"`
}

type DiskItem struct {
	Name              string   `json:"_name"`
	BsdName           string   `json:"bsd_name"`
	DetachableDrive   string   `json:"detachable_drive"`
	DeviceModel       string   `json:"device_model"`
	DeviceRevision    string   `json:"device_revision"`
	DeviceSerial      string   `json:"device_serial"`
	PartitionMapType  string   `json:"partition_map_type"`
	RemovableMedia    string   `json:"removable_media"`
	Size              string   `json:"size"`
	SizeInBytes       int64    `json:"size_in_bytes"`
	SmartStatus       string   `json:"smart_status"`
	SPNVMeLinkSpeed   string   `json:"spnvme_linkspeed"`
	SPNVMeLinkWidth   string   `json:"spnvme_linkwidth"`
	SPNVMeTrimSupport string   `json:"spnvme_trim_support"`
	Volumes           []Volume `json:"volumes"`
}

type SPNVMeDataTypeStruct struct {
	Items []DiskItem `json:"_items"`
	Name  string     `json:"_name"`
}

type SPNVMeDataType struct {
	SPNVMeDataType []SPNVMeDataTypeStruct `json:"SPNVMeDataType"`
}

func getMacInfo[T any]() (result T, fbErr error) {
	defer func() {
		tp := fmt.Sprintf("%T", *new(T))
		if er := recover(); er != nil {
			fbErr = errors.New(fmt.Sprint("Critical error with getting macInfo - структрура - ", tp))
		}
	}()
	tp := fmt.Sprintf("%T", *new(T))
	command := ""
	switch tp {
	case "main.SPAirPortDataType":
		command = "SPAirPortDataType"
	case "main.SPApplicationsDataType":
		command = "SPApplicationsDataType"
	case "main.SPAudioDataType":
		command = "SPAudioDataType"
	case "main.SPDisplaysDataType":
		command = "SPDisplaysDataType"
	case "main.SPHardwareDataType":
		command = "SPHardwareDataType"
	case "main.SPEthernetDataType":
		command = "SPEthernetDataType"
	case "main.SPMemoryDataType":
		command = "SPMemoryDataType"
	case "main.SPNetworkDataType":
		command = "SPNetworkDataType"
	case "main.SPPrintersDataType":
		command = "SPPrintersDataType"
	case "main.SPSerialATADataType":
		command = "SPSerialATADataType"
	case "main.SPSoftwareDataType":
		command = "SPSoftwareDataType"
	case "main.SPNVMeDataType":
		command = "SPNVMeDataType"
	default:
		return result, errors.New(fmt.Sprint("Not implementet for this type - ", tp))
	}
	execCommand := exec.Command("system_profiler", command, "-json")
	var out bytes.Buffer
	execCommand.Stdout = &out
	if err := execCommand.Run(); err != nil {
		return result, err
	} else {
		if err = json.Unmarshal(out.Bytes(), &result); err != nil {
			return result, err
		} else {
			return result, nil
		}
	}
}
