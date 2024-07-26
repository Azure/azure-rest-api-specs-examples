
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.SecurityProfile;
import com.azure.resourcemanager.compute.models.SecurityTypes;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.UefiSettings;
import com.azure.resourcemanager.compute.models.UpgradeMode;
import com.azure.resourcemanager.compute.models.UpgradePolicy;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetManagedDiskParameters;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetStorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMProfile;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithUefiSettings.json
     */
    /**
     * Sample code: Create a scale set with Uefi Settings of secureBoot and vTPM.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAScaleSetWithUefiSettingsOfSecureBootAndVTPM(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets()
            .createOrUpdate("myResourceGroup", "{vmss-name}",
                new VirtualMachineScaleSetInner().withLocation("westus")
                    .withSku(new Sku().withName(
                        "Standard_D2s_v3").withTier(
                            "Standard")
                        .withCapacity(3L))
                    .withUpgradePolicy(new UpgradePolicy().withMode(UpgradeMode.MANUAL))
                    .withVirtualMachineProfile(
                        new VirtualMachineScaleSetVMProfile()
                            .withOsProfile(new VirtualMachineScaleSetOSProfile().withComputerNamePrefix("{vmss-name}")
                                .withAdminUsername("{your-username}").withAdminPassword("fakeTokenPlaceholder"))
                            .withStorageProfile(new VirtualMachineScaleSetStorageProfile()
                                .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                                    .withOffer("windowsserver-gen2preview-preview").withSku("windows10-tvm")
                                    .withVersion("18363.592.2001092016"))
                                .withOsDisk(
                                    new VirtualMachineScaleSetOSDisk().withCaching(CachingTypes.READ_ONLY)
                                        .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                        .withManagedDisk(new VirtualMachineScaleSetManagedDiskParameters()
                                            .withStorageAccountType(StorageAccountTypes.STANDARD_SSD_LRS))))
                            .withNetworkProfile(
                                new VirtualMachineScaleSetNetworkProfile().withNetworkInterfaceConfigurations(
                                    Arrays.asList(new VirtualMachineScaleSetNetworkConfiguration()
                                        .withName("{vmss-name}").withPrimary(true)
                                        .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIpConfiguration()
                                            .withName("{vmss-name}")
                                            .withSubnet(new ApiEntityReference().withId(
                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"))))
                                        .withEnableIpForwarding(true))))
                            .withSecurityProfile(new SecurityProfile()
                                .withUefiSettings(new UefiSettings().withSecureBootEnabled(true).withVTpmEnabled(true))
                                .withSecurityType(SecurityTypes.TRUSTED_LAUNCH)))
                    .withOverprovision(true),
                null, null, com.azure.core.util.Context.NONE);
    }
}
