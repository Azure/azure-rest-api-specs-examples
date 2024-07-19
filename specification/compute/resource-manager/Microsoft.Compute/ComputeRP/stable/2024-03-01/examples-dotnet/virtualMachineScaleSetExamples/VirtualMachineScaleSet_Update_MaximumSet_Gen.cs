using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Update_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineScaleSets_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineScaleSetResource created on azure
// for more information of creating VirtualMachineScaleSetResource, please refer to the document of VirtualMachineScaleSetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string virtualMachineScaleSetName = "aaaaaaaaaaaaa";
ResourceIdentifier virtualMachineScaleSetResourceId = VirtualMachineScaleSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName);
VirtualMachineScaleSetResource virtualMachineScaleSet = client.GetVirtualMachineScaleSetResource(virtualMachineScaleSetResourceId);

// invoke the operation
VirtualMachineScaleSetPatch patch = new VirtualMachineScaleSetPatch()
{
    Sku = new ComputeSku()
    {
        Name = "DSv3-Type1",
        Tier = "aaa",
        Capacity = 7,
    },
    Plan = new ComputePlan()
    {
        Name = "windows2016",
        Publisher = "microsoft-ads",
        Product = "windows-data-science-vm",
        PromotionCode = "aaaaaaaaaa",
    },
    Identity = new ManagedServiceIdentity("SystemAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key3951")] = new UserAssignedIdentity(),
        },
    },
    UpgradePolicy = new VirtualMachineScaleSetUpgradePolicy()
    {
        Mode = VirtualMachineScaleSetUpgradeMode.Manual,
        RollingUpgradePolicy = new RollingUpgradePolicy()
        {
            MaxBatchInstancePercent = 49,
            MaxUnhealthyInstancePercent = 81,
            MaxUnhealthyUpgradedInstancePercent = 98,
            PauseTimeBetweenBatches = "aaaaaaaaaaaaaaa",
            EnableCrossZoneUpgrade = true,
            PrioritizeUnhealthyInstances = true,
            RollbackFailedInstancesOnPolicyBreach = true,
            IsMaxSurgeEnabled = true,
        },
        AutomaticOSUpgradePolicy = new AutomaticOSUpgradePolicy()
        {
            EnableAutomaticOSUpgrade = true,
            DisableAutomaticRollback = true,
            OSRollingUpgradeDeferral = true,
        },
    },
    AutomaticRepairsPolicy = new AutomaticRepairsPolicy()
    {
        Enabled = true,
        GracePeriod = "PT30M",
    },
    VirtualMachineProfile = new VirtualMachineScaleSetUpdateVmProfile()
    {
        OSProfile = new VirtualMachineScaleSetUpdateOSProfile()
        {
            CustomData = "aaaaaaaaaaaaaaaaaaaaaaaaaa",
            WindowsConfiguration = new WindowsConfiguration()
            {
                ProvisionVmAgent = true,
                IsAutomaticUpdatesEnabled = true,
                TimeZone = "aaaaaaaaaaaaaaaa",
                AdditionalUnattendContent =
                {
                new AdditionalUnattendContent()
                {
                PassName = PassName.OobeSystem,
                ComponentName = ComponentName.MicrosoftWindowsShellSetup,
                SettingName = SettingName.AutoLogon,
                Content = "aaaaaaaaaaaaaaaaaaaa",
                }
                },
                PatchSettings = new PatchSettings()
                {
                    PatchMode = WindowsVmGuestPatchMode.AutomaticByPlatform,
                    EnableHotpatching = true,
                    AssessmentMode = WindowsPatchAssessmentMode.ImageDefault,
                    AutomaticByPlatformSettings = new WindowsVmGuestPatchAutomaticByPlatformSettings()
                    {
                        RebootSetting = WindowsVmGuestPatchAutomaticByPlatformRebootSetting.Never,
                    },
                },
                WinRMListeners =
                {
                new WinRMListener()
                {
                Protocol = WinRMListenerProtocolType.Http,
                CertificateUri = new Uri("aaaaaaaaaaaaaaaaaaaaaa"),
                }
                },
            },
            LinuxConfiguration = new LinuxConfiguration()
            {
                IsPasswordAuthenticationDisabled = true,
                SshPublicKeys =
                {
                new SshPublicKeyConfiguration()
                {
                Path = "/home/{your-username}/.ssh/authorized_keys",
                KeyData = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1",
                }
                },
                ProvisionVmAgent = true,
                PatchSettings = new LinuxPatchSettings()
                {
                    PatchMode = LinuxVmGuestPatchMode.ImageDefault,
                    AssessmentMode = LinuxPatchAssessmentMode.ImageDefault,
                },
            },
            Secrets =
            {
            new VaultSecretGroup()
            {
            SourceVaultId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
            VaultCertificates =
            {
            new VaultCertificate()
            {
            CertificateUri = new Uri("aaaaaaa"),
            CertificateStore = "aaaaaaaaaaaaaaaaaaaaaaaaa",
            }
            },
            }
            },
        },
        StorageProfile = new VirtualMachineScaleSetUpdateStorageProfile()
        {
            ImageReference = new ImageReference()
            {
                Publisher = "MicrosoftWindowsServer",
                Offer = "WindowsServer",
                Sku = "2016-Datacenter",
                Version = "latest",
                SharedGalleryImageUniqueId = "aaaaaa",
                Id = new ResourceIdentifier("aaaaaaaaaaaaaaaaaaa"),
            },
            OSDisk = new VirtualMachineScaleSetUpdateOSDisk()
            {
                Caching = CachingType.ReadWrite,
                WriteAcceleratorEnabled = true,
                DiffDiskSettings = new DiffDiskSettings()
                {
                    Option = DiffDiskOption.Local,
                    Placement = DiffDiskPlacement.CacheDisk,
                },
                DiskSizeGB = 6,
                ImageUri = new Uri("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
                VhdContainers =
                {
                "aa"
                },
                ManagedDisk = new VirtualMachineScaleSetManagedDisk()
                {
                    StorageAccountType = StorageAccountType.StandardLrs,
                    DiskEncryptionSetId = new ResourceIdentifier("aaaaaaaaaaaa"),
                },
            },
            DataDisks =
            {
            new VirtualMachineScaleSetDataDisk(26,DiskCreateOptionType.Empty)
            {
            Name = "aaaaaaaaaaaaaaaaaaaaaaaaaa",
            Caching = CachingType.None,
            WriteAcceleratorEnabled = true,
            DiskSizeGB = 1023,
            ManagedDisk = new VirtualMachineScaleSetManagedDisk()
            {
            StorageAccountType = StorageAccountType.StandardLrs,
            DiskEncryptionSetId = new ResourceIdentifier("aaaaaaaaaaaa"),
            },
            DiskIopsReadWrite = 28,
            DiskMBpsReadWrite = 15,
            }
            },
        },
        NetworkProfile = new VirtualMachineScaleSetUpdateNetworkProfile()
        {
            HealthProbeId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
            NetworkInterfaceConfigurations =
            {
            new VirtualMachineScaleSetUpdateNetworkConfiguration()
            {
            Name = "aaaaaaaa",
            Primary = true,
            EnableAcceleratedNetworking = true,
            EnableFpga = true,
            NetworkSecurityGroupId = new ResourceIdentifier("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
            DnsServers =
            {
            },
            IPConfigurations =
            {
            new VirtualMachineScaleSetUpdateIPConfiguration()
            {
            Name = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            SubnetId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
            Primary = true,
            PublicIPAddressConfiguration = new VirtualMachineScaleSetUpdatePublicIPAddressConfiguration()
            {
            Name = "a",
            IdleTimeoutInMinutes = 3,
            DnsSettings = new VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings("aaaaaaaaaaaaaaaaaa"),
            DeleteOption = ComputeDeleteOption.Delete,
            },
            PrivateIPAddressVersion = IPVersion.IPv4,
            ApplicationGatewayBackendAddressPools =
            {
            new WritableSubResource()
            {
            Id = new ResourceIdentifier("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
            }
            },
            ApplicationSecurityGroups =
            {
            new WritableSubResource()
            {
            Id = new ResourceIdentifier("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
            }
            },
            LoadBalancerBackendAddressPools =
            {
            new WritableSubResource()
            {
            Id = new ResourceIdentifier("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
            }
            },
            LoadBalancerInboundNatPools =
            {
            new WritableSubResource()
            {
            Id = new ResourceIdentifier("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
            }
            },
            }
            },
            EnableIPForwarding = true,
            DeleteOption = ComputeDeleteOption.Delete,
            }
            },
            NetworkApiVersion = NetworkApiVersion.TwoThousandTwenty1101,
        },
        SecurityProfile = new SecurityProfile()
        {
            UefiSettings = new UefiSettings()
            {
                IsSecureBootEnabled = true,
                IsVirtualTpmEnabled = true,
            },
            EncryptionAtHost = true,
            SecurityType = SecurityType.TrustedLaunch,
        },
        BootDiagnostics = new BootDiagnostics()
        {
            Enabled = true,
            StorageUri = new Uri("http://{existing-storage-account-name}.blob.core.windows.net"),
        },
        ExtensionProfile = new VirtualMachineScaleSetExtensionProfile()
        {
            Extensions =
            {
            new VirtualMachineScaleSetExtensionData()
            {
            ForceUpdateTag = "aaaaaaaaa",
            Publisher = "{extension-Publisher}",
            ExtensionType = "{extension-Type}",
            TypeHandlerVersion = "{handler-version}",
            AutoUpgradeMinorVersion = true,
            EnableAutomaticUpgrade = true,
            Settings = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
            {
            }),
            ProtectedSettings = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
            {
            }),
            ProvisionAfterExtensions =
            {
            "aa"
            },
            SuppressFailures = true,
            }
            },
            ExtensionsTimeBudget = "PT1H20M",
        },
        LicenseType = "aaaaaaaaaaaa",
        BillingMaxPrice = -1,
        ScheduledEventsProfile = new ComputeScheduledEventsProfile()
        {
            TerminateNotificationProfile = new TerminateNotificationProfile()
            {
                NotBeforeTimeout = "PT10M",
                Enable = true,
            },
        },
        UserData = "aaaaaaaaaaaaa",
    },
    Overprovision = true,
    DoNotRunExtensionsOnOverprovisionedVms = true,
    SinglePlacementGroup = true,
    AdditionalCapabilities = new AdditionalCapabilities()
    {
        UltraSsdEnabled = true,
        HibernationEnabled = true,
    },
    ScaleInPolicy = new ScaleInPolicy()
    {
        Rules =
        {
        VirtualMachineScaleSetScaleInRule.OldestVm
        },
        ForceDeletion = true,
    },
    ProximityPlacementGroupId = new ResourceIdentifier("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
    Tags =
    {
    ["key246"] = "aaaaaaaaaaaaaaaaaaaaaaaa",
    },
};
ArmOperation<VirtualMachineScaleSetResource> lro = await virtualMachineScaleSet.UpdateAsync(WaitUntil.Completed, patch);
VirtualMachineScaleSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineScaleSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
