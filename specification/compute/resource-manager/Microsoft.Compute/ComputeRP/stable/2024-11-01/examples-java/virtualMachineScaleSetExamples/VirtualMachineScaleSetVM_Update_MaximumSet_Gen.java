
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetVMInner;
import com.azure.resourcemanager.compute.models.AdditionalCapabilities;
import com.azure.resourcemanager.compute.models.AdditionalUnattendContent;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.BootDiagnostics;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.ComponentNames;
import com.azure.resourcemanager.compute.models.DataDisk;
import com.azure.resourcemanager.compute.models.DeleteOptions;
import com.azure.resourcemanager.compute.models.DiagnosticsProfile;
import com.azure.resourcemanager.compute.models.DiffDiskOptions;
import com.azure.resourcemanager.compute.models.DiffDiskPlacement;
import com.azure.resourcemanager.compute.models.DiffDiskSettings;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.DiskDeleteOptionTypes;
import com.azure.resourcemanager.compute.models.DiskDetachOptionTypes;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetParameters;
import com.azure.resourcemanager.compute.models.DiskEncryptionSettings;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.IpVersion;
import com.azure.resourcemanager.compute.models.IpVersions;
import com.azure.resourcemanager.compute.models.KeyVaultKeyReference;
import com.azure.resourcemanager.compute.models.KeyVaultSecretReference;
import com.azure.resourcemanager.compute.models.LinuxConfiguration;
import com.azure.resourcemanager.compute.models.LinuxPatchAssessmentMode;
import com.azure.resourcemanager.compute.models.LinuxPatchSettings;
import com.azure.resourcemanager.compute.models.LinuxVMGuestPatchMode;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkApiVersion;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;
import com.azure.resourcemanager.compute.models.PassNames;
import com.azure.resourcemanager.compute.models.PatchSettings;
import com.azure.resourcemanager.compute.models.Plan;
import com.azure.resourcemanager.compute.models.ProtocolTypes;
import com.azure.resourcemanager.compute.models.PublicIpAddressSku;
import com.azure.resourcemanager.compute.models.PublicIpAddressSkuName;
import com.azure.resourcemanager.compute.models.PublicIpAddressSkuTier;
import com.azure.resourcemanager.compute.models.PublicIpAllocationMethod;
import com.azure.resourcemanager.compute.models.SecurityProfile;
import com.azure.resourcemanager.compute.models.SecurityTypes;
import com.azure.resourcemanager.compute.models.SettingNames;
import com.azure.resourcemanager.compute.models.SshConfiguration;
import com.azure.resourcemanager.compute.models.SshPublicKey;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.UefiSettings;
import com.azure.resourcemanager.compute.models.VMSizeProperties;
import com.azure.resourcemanager.compute.models.VirtualHardDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineIpTag;
import com.azure.resourcemanager.compute.models.VirtualMachineNetworkInterfaceConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineNetworkInterfaceDnsSettingsConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineNetworkInterfaceIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachinePublicIpAddressConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachinePublicIpAddressDnsSettingsConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIpTag;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkConfigurationDnsSettings;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetPublicIpAddressConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetPublicIpAddressConfigurationDnsSettings;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMNetworkProfileConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMProtectionPolicy;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import com.azure.resourcemanager.compute.models.WinRMConfiguration;
import com.azure.resourcemanager.compute.models.WinRMListener;
import com.azure.resourcemanager.compute.models.WindowsConfiguration;
import com.azure.resourcemanager.compute.models.WindowsPatchAssessmentMode;
import com.azure.resourcemanager.compute.models.WindowsVMGuestPatchMode;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachineScaleSetVMs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Update_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetVMUpdateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().update("rgcompute",
            "aaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            new VirtualMachineScaleSetVMInner().withLocation("westus").withTags(mapOf())
                .withPlan(new Plan().withName("aaaaaaaaaa").withPublisher("aaaaaaaaaaaaaaaaaaaaaa")
                    .withProduct("aaaaaaaaaaaaaaaaaaaa").withPromotionCode("fakeTokenPlaceholder"))
                .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.BASIC_A0)
                    .withVmSizeProperties(new VMSizeProperties().withVCpusAvailable(9).withVCpusPerCore(12)))
                .withStorageProfile(new StorageProfile()
                    .withImageReference(new ImageReference().withId("a").withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2012-R2-Datacenter").withVersion("4.127.20180315")
                        .withSharedGalleryImageId("aaaaaaaaaaaaaaaaaaaa"))
                    .withOsDisk(new OSDisk().withOsType(OperatingSystemTypes.WINDOWS).withEncryptionSettings(
                        new DiskEncryptionSettings().withDiskEncryptionKey(new KeyVaultSecretReference()
                            .withSecretUrl("fakeTokenPlaceholder")
                            .withSourceVault(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))
                            .withKeyEncryptionKey(new KeyVaultKeyReference().withKeyUrl("fakeTokenPlaceholder")
                                .withSourceVault(new SubResource().withId(
                                    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))
                            .withEnabled(true))
                        .withName("vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc")
                        .withVhd(new VirtualHardDisk().withUri(
                            "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"))
                        .withImage(new VirtualHardDisk().withUri(
                            "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"))
                        .withCaching(CachingTypes.NONE).withWriteAcceleratorEnabled(true)
                        .withDiffDiskSettings(new DiffDiskSettings().withOption(DiffDiskOptions.LOCAL)
                            .withPlacement(DiffDiskPlacement.CACHE_DISK))
                        .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE).withDiskSizeGB(127)
                        .withManagedDisk(new ManagedDiskParameters().withId(
                            "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc")
                            .withStorageAccountType(StorageAccountTypes.STANDARD_LRS)
                            .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId("aaaaaaaaaaaa")))
                        .withDeleteOption(DiskDeleteOptionTypes.DELETE))
                    .withDataDisks(Arrays.asList(new DataDisk().withLun(1)
                        .withName("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d")
                        .withVhd(new VirtualHardDisk().withUri(
                            "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"))
                        .withImage(new VirtualHardDisk().withUri(
                            "https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"))
                        .withCaching(CachingTypes.NONE).withWriteAcceleratorEnabled(true)
                        .withCreateOption(DiskCreateOptionTypes.EMPTY).withDiskSizeGB(128)
                        .withManagedDisk(new ManagedDiskParameters().withId(
                            "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d")
                            .withStorageAccountType(StorageAccountTypes.STANDARD_LRS)
                            .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId("aaaaaaaaaaaa")))
                        .withToBeDetached(true).withDetachOption(DiskDetachOptionTypes.FORCE_DETACH)
                        .withDeleteOption(DiskDeleteOptionTypes.DELETE))))
                .withAdditionalCapabilities(
                    new AdditionalCapabilities().withUltraSsdEnabled(true).withHibernationEnabled(true))
                .withOsProfile(new OSProfile().withComputerName("test000000").withAdminUsername("Foo12")
                    .withAdminPassword("fakeTokenPlaceholder").withCustomData("aaaa")
                    .withWindowsConfiguration(new WindowsConfiguration().withProvisionVMAgent(true)
                        .withEnableAutomaticUpdates(true).withTimeZone("aaaaaaaaaaaaaaaaaaaaaaaaaaa")
                        .withAdditionalUnattendContent(
                            Arrays.asList(new AdditionalUnattendContent().withPassName(PassNames.OOBE_SYSTEM)
                                .withComponentName(ComponentNames.MICROSOFT_WINDOWS_SHELL_SETUP)
                                .withSettingName(SettingNames.AUTO_LOGON).withContent("aaaaaaaaaaaaaaaaaaaa")))
                        .withPatchSettings(new PatchSettings().withPatchMode(WindowsVMGuestPatchMode.MANUAL)
                            .withEnableHotpatching(true).withAssessmentMode(WindowsPatchAssessmentMode.IMAGE_DEFAULT))
                        .withWinRM(new WinRMConfiguration().withListeners(Arrays.asList(new WinRMListener()
                            .withProtocol(ProtocolTypes.HTTP).withCertificateUrl("aaaaaaaaaaaaaaaaaaaaaa")))))
                    .withLinuxConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(true)
                        .withSsh(new SshConfiguration().withPublicKeys(
                            Arrays.asList(new SshPublicKey().withPath("aaa").withKeyData("fakeTokenPlaceholder"))))
                        .withProvisionVMAgent(true)
                        .withPatchSettings(new LinuxPatchSettings().withPatchMode(LinuxVMGuestPatchMode.IMAGE_DEFAULT)
                            .withAssessmentMode(LinuxPatchAssessmentMode.IMAGE_DEFAULT)))
                    .withSecrets(Arrays.asList()).withAllowExtensionOperations(true)
                    .withRequireGuestProvisionSignal(true))
                .withSecurityProfile(new SecurityProfile()
                    .withUefiSettings(new UefiSettings().withSecureBootEnabled(true).withVTpmEnabled(true))
                    .withEncryptionAtHost(true).withSecurityType(SecurityTypes.TRUSTED_LAUNCH))
                .withNetworkProfile(new NetworkProfile()
                    .withNetworkInterfaces(Arrays.asList(new NetworkInterfaceReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415")
                        .withPrimary(true).withDeleteOption(DeleteOptions.DELETE)))
                    .withNetworkApiVersion(NetworkApiVersion.TWO_ZERO_TWO_ZERO_ONE_ONE_ZERO_ONE)
                    .withNetworkInterfaceConfigurations(Arrays.asList(new VirtualMachineNetworkInterfaceConfiguration()
                        .withName("aaaaaaaaaaa").withPrimary(true).withDeleteOption(DeleteOptions.DELETE)
                        .withEnableAcceleratedNetworking(true).withEnableFpga(true).withEnableIpForwarding(true)
                        .withNetworkSecurityGroup(new SubResource().withId(
                            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"))
                        .withDnsSettings(new VirtualMachineNetworkInterfaceDnsSettingsConfiguration()
                            .withDnsServers(Arrays.asList("aaaaaa")))
                        .withIpConfigurations(Arrays.asList(new VirtualMachineNetworkInterfaceIpConfiguration()
                            .withName("aa")
                            .withSubnet(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"))
                            .withPrimary(true)
                            .withPublicIpAddressConfiguration(new VirtualMachinePublicIpAddressConfiguration()
                                .withName("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
                                .withSku(new PublicIpAddressSku().withName(PublicIpAddressSkuName.BASIC)
                                    .withTier(PublicIpAddressSkuTier.REGIONAL))
                                .withIdleTimeoutInMinutes(2).withDeleteOption(DeleteOptions.DELETE)
                                .withDnsSettings(new VirtualMachinePublicIpAddressDnsSettingsConfiguration()
                                    .withDomainNameLabel("aaaaaaaaaaaaaaaaaaaaaaaaa"))
                                .withIpTags(Arrays.asList(new VirtualMachineIpTag()
                                    .withIpTagType("aaaaaaaaaaaaaaaaaaaaaaaaa").withTag("aaaaaaaaaaaaaaaaaaaa")))
                                .withPublicIpPrefix(new SubResource().withId(
                                    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"))
                                .withPublicIpAddressVersion(IpVersions.IPV4)
                                .withPublicIpAllocationMethod(PublicIpAllocationMethod.DYNAMIC))
                            .withPrivateIpAddressVersion(IpVersions.IPV4)
                            .withApplicationSecurityGroups(Arrays.asList(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))
                            .withApplicationGatewayBackendAddressPools(Arrays.asList(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))
                            .withLoadBalancerBackendAddressPools(Arrays.asList(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))))
                        .withDscpConfiguration(new SubResource().withId(
                            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))))
                .withNetworkProfileConfiguration(new VirtualMachineScaleSetVMNetworkProfileConfiguration()
                    .withNetworkInterfaceConfigurations(Arrays.asList(new VirtualMachineScaleSetNetworkConfiguration()
                        .withName("vmsstestnetconfig5415").withPrimary(true).withEnableAcceleratedNetworking(true)
                        .withEnableFpga(true)
                        .withNetworkSecurityGroup(new SubResource().withId(
                            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"))
                        .withDnsSettings(
                            new VirtualMachineScaleSetNetworkConfigurationDnsSettings().withDnsServers(Arrays.asList()))
                        .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIpConfiguration()
                            .withName("vmsstestnetconfig9693")
                            .withSubnet(new ApiEntityReference().withId(
                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vn4071/subnets/sn5503"))
                            .withPrimary(true)
                            .withPublicIpAddressConfiguration(new VirtualMachineScaleSetPublicIpAddressConfiguration()
                                .withName("aaaaaaaaaaaaaaaaaa")
                                .withSku(new PublicIpAddressSku()
                                    .withName(PublicIpAddressSkuName.BASIC).withTier(PublicIpAddressSkuTier.REGIONAL))
                                .withIdleTimeoutInMinutes(18)
                                .withDnsSettings(new VirtualMachineScaleSetPublicIpAddressConfigurationDnsSettings()
                                    .withDomainNameLabel("aaaaaaaaaaaaaaaaaa"))
                                .withIpTags(Arrays.asList(new VirtualMachineScaleSetIpTag().withIpTagType("aaaaaaa")
                                    .withTag("aaaaaaaaaaaaaaaaaaaaaaaaaaa")))
                                .withPublicIpPrefix(new SubResource().withId(
                                    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"))
                                .withPublicIpAddressVersion(IpVersion.IPV4).withDeleteOption(DeleteOptions.DELETE))
                            .withPrivateIpAddressVersion(IpVersion.IPV4)
                            .withApplicationGatewayBackendAddressPools(Arrays.asList(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))
                            .withApplicationSecurityGroups(Arrays.asList(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))
                            .withLoadBalancerBackendAddressPools(Arrays.asList(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))
                            .withLoadBalancerInboundNatPools(Arrays.asList(new SubResource().withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}")))))
                        .withEnableIpForwarding(true).withDeleteOption(DeleteOptions.DELETE))))
                .withDiagnosticsProfile(new DiagnosticsProfile()
                    .withBootDiagnostics(new BootDiagnostics().withEnabled(true).withStorageUri("aaaaaaaaaaaaa")))
                .withAvailabilitySet(new SubResource().withId(
                    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"))
                .withLicenseType("aaaaaaaaaa")
                .withProtectionPolicy(new VirtualMachineScaleSetVMProtectionPolicy().withProtectFromScaleIn(true)
                    .withProtectFromScaleSetActions(true))
                .withUserData("RXhhbXBsZSBVc2VyRGF0YQ=="),
            null, null, com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
