
import com.azure.core.management.SubResource;
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetExtensionInner;
import com.azure.resourcemanager.compute.models.AdditionalCapabilities;
import com.azure.resourcemanager.compute.models.AdditionalUnattendContent;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.AutomaticOSUpgradePolicy;
import com.azure.resourcemanager.compute.models.AutomaticRepairsPolicy;
import com.azure.resourcemanager.compute.models.BillingProfile;
import com.azure.resourcemanager.compute.models.BootDiagnostics;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.ComponentNames;
import com.azure.resourcemanager.compute.models.DeleteOptions;
import com.azure.resourcemanager.compute.models.DiagnosticsProfile;
import com.azure.resourcemanager.compute.models.DiffDiskOptions;
import com.azure.resourcemanager.compute.models.DiffDiskPlacement;
import com.azure.resourcemanager.compute.models.DiffDiskSettings;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetParameters;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.IpVersion;
import com.azure.resourcemanager.compute.models.LinuxConfiguration;
import com.azure.resourcemanager.compute.models.LinuxPatchAssessmentMode;
import com.azure.resourcemanager.compute.models.LinuxPatchSettings;
import com.azure.resourcemanager.compute.models.LinuxVMGuestPatchMode;
import com.azure.resourcemanager.compute.models.NetworkApiVersion;
import com.azure.resourcemanager.compute.models.PassNames;
import com.azure.resourcemanager.compute.models.PatchSettings;
import com.azure.resourcemanager.compute.models.Plan;
import com.azure.resourcemanager.compute.models.ProtocolTypes;
import com.azure.resourcemanager.compute.models.ResourceIdentityType;
import com.azure.resourcemanager.compute.models.RollingUpgradePolicy;
import com.azure.resourcemanager.compute.models.ScaleInPolicy;
import com.azure.resourcemanager.compute.models.ScheduledEventsProfile;
import com.azure.resourcemanager.compute.models.SecurityProfile;
import com.azure.resourcemanager.compute.models.SecurityTypes;
import com.azure.resourcemanager.compute.models.SettingNames;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.SshConfiguration;
import com.azure.resourcemanager.compute.models.SshPublicKey;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.TerminateNotificationProfile;
import com.azure.resourcemanager.compute.models.UefiSettings;
import com.azure.resourcemanager.compute.models.UpgradeMode;
import com.azure.resourcemanager.compute.models.UpgradePolicy;
import com.azure.resourcemanager.compute.models.VaultCertificate;
import com.azure.resourcemanager.compute.models.VaultSecretGroup;
import com.azure.resourcemanager.compute.models.VirtualHardDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineIdentityUserAssignedIdentities;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetDataDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetExtensionProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIdentity;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetManagedDiskParameters;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkConfigurationDnsSettings;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetPublicIpAddressConfigurationDnsSettings;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetScaleInRules;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdate;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdateIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdateNetworkConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdateNetworkProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdateOSDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdateOSProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdatePublicIpAddressConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdateStorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetUpdateVMProfile;
import com.azure.resourcemanager.compute.models.WinRMConfiguration;
import com.azure.resourcemanager.compute.models.WinRMListener;
import com.azure.resourcemanager.compute.models.WindowsConfiguration;
import com.azure.resourcemanager.compute.models.WindowsPatchAssessmentMode;
import com.azure.resourcemanager.compute.models.WindowsVMGuestPatchAutomaticByPlatformRebootSetting;
import com.azure.resourcemanager.compute.models.WindowsVMGuestPatchAutomaticByPlatformSettings;
import com.azure.resourcemanager.compute.models.WindowsVMGuestPatchMode;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachineScaleSets Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Update_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetUpdateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().update("rgcompute",
            "aaaaaaaaaaaaa",
            new VirtualMachineScaleSetUpdate().withTags(mapOf("key246", "fakeTokenPlaceholder"))
                .withSku(new Sku().withName("DSv3-Type1").withTier("aaa").withCapacity(7L))
                .withPlan(new Plan().withName("windows2016").withPublisher("microsoft-ads")
                    .withProduct("windows-data-science-vm").withPromotionCode("fakeTokenPlaceholder"))
                .withIdentity(new VirtualMachineScaleSetIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)
                    .withUserAssignedIdentities(mapOf("key3951", new VirtualMachineIdentityUserAssignedIdentities())))
                .withZones(Arrays.asList("1", "2", "3"))
                .withUpgradePolicy(
                    new UpgradePolicy().withMode(UpgradeMode.MANUAL)
                        .withRollingUpgradePolicy(new RollingUpgradePolicy()
                            .withMaxBatchInstancePercent(49).withMaxUnhealthyInstancePercent(81)
                            .withMaxUnhealthyUpgradedInstancePercent(98).withPauseTimeBetweenBatches(
                                "aaaaaaaaaaaaaaa")
                            .withEnableCrossZoneUpgrade(true).withPrioritizeUnhealthyInstances(true)
                            .withRollbackFailedInstancesOnPolicyBreach(true).withMaxSurge(true))
                        .withAutomaticOSUpgradePolicy(new AutomaticOSUpgradePolicy()
                            .withEnableAutomaticOSUpgrade(true).withDisableAutomaticRollback(
                                true)
                            .withOsRollingUpgradeDeferral(true)))
                .withAutomaticRepairsPolicy(new AutomaticRepairsPolicy().withEnabled(true).withGracePeriod("PT30M"))
                .withVirtualMachineProfile(
                    new VirtualMachineScaleSetUpdateVMProfile()
                        .withOsProfile(
                            new VirtualMachineScaleSetUpdateOSProfile().withCustomData("aaaaaaaaaaaaaaaaaaaaaaaaaa")
                                .withWindowsConfiguration(new WindowsConfiguration()
                                    .withProvisionVMAgent(true).withEnableAutomaticUpdates(true)
                                    .withTimeZone("aaaaaaaaaaaaaaaa")
                                    .withAdditionalUnattendContent(Arrays
                                        .asList(new AdditionalUnattendContent().withPassName(PassNames.OOBE_SYSTEM)
                                            .withComponentName(ComponentNames.MICROSOFT_WINDOWS_SHELL_SETUP)
                                            .withSettingName(SettingNames.AUTO_LOGON)
                                            .withContent("aaaaaaaaaaaaaaaaaaaa")))
                                    .withPatchSettings(
                                        new PatchSettings().withPatchMode(WindowsVMGuestPatchMode.AUTOMATIC_BY_PLATFORM)
                                            .withEnableHotpatching(true)
                                            .withAssessmentMode(WindowsPatchAssessmentMode.IMAGE_DEFAULT)
                                            .withAutomaticByPlatformSettings(
                                                new WindowsVMGuestPatchAutomaticByPlatformSettings().withRebootSetting(
                                                    WindowsVMGuestPatchAutomaticByPlatformRebootSetting.NEVER)))
                                    .withWinRM(new WinRMConfiguration().withListeners(
                                        Arrays.asList(new WinRMListener().withProtocol(ProtocolTypes.HTTP)
                                            .withCertificateUrl("aaaaaaaaaaaaaaaaaaaaaa")))))
                                .withLinuxConfiguration(
                                    new LinuxConfiguration().withDisablePasswordAuthentication(true)
                                        .withSsh(new SshConfiguration().withPublicKeys(Arrays.asList(
                                            new SshPublicKey().withPath("/home/{your-username}/.ssh/authorized_keys")
                                                .withKeyData("fakeTokenPlaceholder"))))
                                        .withProvisionVMAgent(true).withPatchSettings(
                                            new LinuxPatchSettings()
                                                .withPatchMode(LinuxVMGuestPatchMode.IMAGE_DEFAULT)
                                                .withAssessmentMode(LinuxPatchAssessmentMode.IMAGE_DEFAULT)))
                                .withSecrets(Arrays.asList(new VaultSecretGroup()
                                    .withSourceVault(new SubResource().withId(
                                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"))
                                    .withVaultCertificates(
                                        Arrays.asList(new VaultCertificate().withCertificateUrl("aaaaaaa")
                                            .withCertificateStore("aaaaaaaaaaaaaaaaaaaaaaaaa"))))))
                        .withStorageProfile(new VirtualMachineScaleSetUpdateStorageProfile()
                            .withImageReference(new ImageReference()
                                .withId("aaaaaaaaaaaaaaaaaaa").withPublisher("MicrosoftWindowsServer")
                                .withOffer("WindowsServer").withSku("2016-Datacenter").withVersion("latest")
                                .withSharedGalleryImageId("aaaaaa"))
                            .withOsDisk(new VirtualMachineScaleSetUpdateOSDisk().withCaching(CachingTypes.READ_WRITE)
                                .withWriteAcceleratorEnabled(true)
                                .withDiffDiskSettings(new DiffDiskSettings()
                                    .withOption(DiffDiskOptions.LOCAL).withPlacement(DiffDiskPlacement.CACHE_DISK))
                                .withDiskSizeGB(6)
                                .withImage(new VirtualHardDisk().withUri(
                                    "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"))
                                .withVhdContainers(Arrays.asList("aa"))
                                .withManagedDisk(new VirtualMachineScaleSetManagedDiskParameters()
                                    .withStorageAccountType(StorageAccountTypes.STANDARD_LRS)
                                    .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId("aaaaaaaaaaaa"))))
                            .withDataDisks(Arrays.asList(new VirtualMachineScaleSetDataDisk()
                                .withName("aaaaaaaaaaaaaaaaaaaaaaaaaa").withLun(26).withCaching(CachingTypes.NONE)
                                .withWriteAcceleratorEnabled(true).withCreateOption(DiskCreateOptionTypes.EMPTY)
                                .withDiskSizeGB(1023)
                                .withManagedDisk(new VirtualMachineScaleSetManagedDiskParameters()
                                    .withStorageAccountType(StorageAccountTypes.STANDARD_LRS)
                                    .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId("aaaaaaaaaaaa")))
                                .withDiskIopsReadWrite(28L).withDiskMBpsReadWrite(15L))))
                        .withNetworkProfile(new VirtualMachineScaleSetUpdateNetworkProfile()
                            .withHealthProbe(new ApiEntityReference().withId(
                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"))
                            .withNetworkInterfaceConfigurations(Arrays
                                .asList(new VirtualMachineScaleSetUpdateNetworkConfiguration().withName("aaaaaaaa")
                                    .withPrimary(true).withEnableAcceleratedNetworking(true).withEnableFpga(true)
                                    .withNetworkSecurityGroup(
                                        new SubResource().withId(
                                            "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"))
                                    .withDnsSettings(new VirtualMachineScaleSetNetworkConfigurationDnsSettings()
                                        .withDnsServers(Arrays.asList()))
                                    .withIpConfigurations(Arrays.asList(
                                        new VirtualMachineScaleSetUpdateIpConfiguration().withName(
                                            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
                                            .withSubnet(new ApiEntityReference()
                                                .withId(
                                                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"))
                                            .withPrimary(true).withPublicIpAddressConfiguration(
                                                new VirtualMachineScaleSetUpdatePublicIpAddressConfiguration().withName(
                                                    "a").withIdleTimeoutInMinutes(3).withDnsSettings(
                                                        new VirtualMachineScaleSetPublicIpAddressConfigurationDnsSettings()
                                                            .withDomainNameLabel("aaaaaaaaaaaaaaaaaa"))
                                                    .withDeleteOption(DeleteOptions.DELETE))
                                            .withPrivateIpAddressVersion(IpVersion.IPV4)
                                            .withApplicationGatewayBackendAddressPools(
                                                Arrays.asList(new SubResource().withId(
                                                    "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")))
                                            .withApplicationSecurityGroups(Arrays.asList(new SubResource().withId(
                                                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")))
                                            .withLoadBalancerBackendAddressPools(Arrays.asList(new SubResource().withId(
                                                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")))
                                            .withLoadBalancerInboundNatPools(Arrays.asList(new SubResource().withId(
                                                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")))))
                                    .withEnableIpForwarding(true).withDeleteOption(DeleteOptions.DELETE)))
                            .withNetworkApiVersion(NetworkApiVersion.TWO_ZERO_TWO_ZERO_ONE_ONE_ZERO_ONE))
                        .withSecurityProfile(new SecurityProfile()
                            .withUefiSettings(new UefiSettings().withSecureBootEnabled(true).withVTpmEnabled(true))
                            .withEncryptionAtHost(true).withSecurityType(SecurityTypes.TRUSTED_LAUNCH))
                        .withDiagnosticsProfile(
                            new DiagnosticsProfile().withBootDiagnostics(new BootDiagnostics().withEnabled(true)
                                .withStorageUri("http://{existing-storage-account-name}.blob.core.windows.net")))
                        .withExtensionProfile(new VirtualMachineScaleSetExtensionProfile()
                            .withExtensions(Arrays.asList(new VirtualMachineScaleSetExtensionInner()
                                .withName("{extension-name}").withForceUpdateTag("aaaaaaaaa")
                                .withPublisher("{extension-Publisher}").withTypePropertiesType("{extension-Type}")
                                .withTypeHandlerVersion("{handler-version}").withAutoUpgradeMinorVersion(true)
                                .withEnableAutomaticUpgrade(true)
                                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                                    .deserialize("{}", Object.class, SerializerEncoding.JSON))
                                .withProtectedSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                                    .deserialize("{}", Object.class, SerializerEncoding.JSON))
                                .withProvisionAfterExtensions(Arrays.asList("aa")).withSuppressFailures(true)))
                            .withExtensionsTimeBudget("PT1H20M"))
                        .withLicenseType("aaaaaaaaaaaa").withBillingProfile(new BillingProfile().withMaxPrice(-1.0D))
                        .withScheduledEventsProfile(new ScheduledEventsProfile().withTerminateNotificationProfile(
                            new TerminateNotificationProfile().withNotBeforeTimeout("PT10M").withEnable(true)))
                        .withUserData("aaaaaaaaaaaaa"))
                .withOverprovision(true).withDoNotRunExtensionsOnOverprovisionedVMs(true).withSinglePlacementGroup(true)
                .withAdditionalCapabilities(
                    new AdditionalCapabilities().withUltraSsdEnabled(true).withHibernationEnabled(true))
                .withScaleInPolicy(new ScaleInPolicy()
                    .withRules(Arrays.asList(VirtualMachineScaleSetScaleInRules.OLDEST_VM)).withForceDeletion(true))
                .withProximityPlacementGroup(new SubResource().withId(
                    "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot")),
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
