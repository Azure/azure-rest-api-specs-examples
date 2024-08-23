
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.AutomaticOSUpgradePolicy;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ServiceArtifactReference;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.UpgradeMode;
import com.azure.resourcemanager.compute.models.UpgradePolicy;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIpConfiguration;
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
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithServiceArtifactReference.json
     */
    /**
     * Sample code: Create a scale set with Service Artifact Reference.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAScaleSetWithServiceArtifactReference(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().createOrUpdate("myResourceGroup",
            "{vmss-name}",
            new VirtualMachineScaleSetInner().withLocation("eastus2euap")
                .withSku(new Sku().withName("Standard_A1").withTier("Standard").withCapacity(3L))
                .withUpgradePolicy(new UpgradePolicy().withMode(UpgradeMode.AUTOMATIC)
                    .withAutomaticOSUpgradePolicy(new AutomaticOSUpgradePolicy().withEnableAutomaticOSUpgrade(true)))
                .withVirtualMachineProfile(new VirtualMachineScaleSetVMProfile()
                    .withOsProfile(new VirtualMachineScaleSetOSProfile().withComputerNamePrefix("{vmss-name}")
                        .withAdminUsername("{your-username}").withAdminPassword("fakeTokenPlaceholder"))
                    .withStorageProfile(new VirtualMachineScaleSetStorageProfile()
                        .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                            .withOffer("WindowsServer").withSku("2022-Datacenter").withVersion("latest"))
                        .withOsDisk(new VirtualMachineScaleSetOSDisk().withName("osDisk")
                            .withCaching(CachingTypes.READ_WRITE).withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)))
                    .withNetworkProfile(new VirtualMachineScaleSetNetworkProfile().withNetworkInterfaceConfigurations(
                        Arrays.asList(new VirtualMachineScaleSetNetworkConfiguration().withName("{vmss-name}")
                            .withPrimary(true)
                            .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIpConfiguration()
                                .withName("{vmss-name}")
                                .withSubnet(new ApiEntityReference().withId(
                                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"))))
                            .withEnableIpForwarding(true))))
                    .withServiceArtifactReference(new ServiceArtifactReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/serviceArtifacts/serviceArtifactName/vmArtifactsProfiles/vmArtifactsProfilesName")))
                .withOverprovision(true),
            null, null, com.azure.core.util.Context.NONE);
    }
}
