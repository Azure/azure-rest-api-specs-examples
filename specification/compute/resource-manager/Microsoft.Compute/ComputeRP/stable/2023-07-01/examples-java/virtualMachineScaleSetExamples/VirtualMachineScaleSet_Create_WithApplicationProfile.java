import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.ApplicationProfile;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.UpgradeMode;
import com.azure.resourcemanager.compute.models.UpgradePolicy;
import com.azure.resourcemanager.compute.models.VMGalleryApplication;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetManagedDiskParameters;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetStorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMProfile;
import java.util.Arrays;

/** Samples for VirtualMachineScaleSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithApplicationProfile.json
     */
    /**
     * Sample code: Create a scale set with Application Profile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAScaleSetWithApplicationProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .createOrUpdate(
                "myResourceGroup",
                "{vmss-name}",
                new VirtualMachineScaleSetInner()
                    .withLocation("westus")
                    .withSku(new Sku().withName("Standard_D1_v2").withTier("Standard").withCapacity(3L))
                    .withUpgradePolicy(new UpgradePolicy().withMode(UpgradeMode.MANUAL))
                    .withVirtualMachineProfile(
                        new VirtualMachineScaleSetVMProfile()
                            .withOsProfile(
                                new VirtualMachineScaleSetOSProfile()
                                    .withComputerNamePrefix("{vmss-name}")
                                    .withAdminUsername("{your-username}")
                                    .withAdminPassword("fakeTokenPlaceholder"))
                            .withStorageProfile(
                                new VirtualMachineScaleSetStorageProfile()
                                    .withImageReference(
                                        new ImageReference()
                                            .withPublisher("MicrosoftWindowsServer")
                                            .withOffer("WindowsServer")
                                            .withSku("2016-Datacenter")
                                            .withVersion("latest"))
                                    .withOsDisk(
                                        new VirtualMachineScaleSetOSDisk()
                                            .withCaching(CachingTypes.READ_WRITE)
                                            .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                            .withManagedDisk(
                                                new VirtualMachineScaleSetManagedDiskParameters()
                                                    .withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                            .withNetworkProfile(
                                new VirtualMachineScaleSetNetworkProfile()
                                    .withNetworkInterfaceConfigurations(
                                        Arrays
                                            .asList(
                                                new VirtualMachineScaleSetNetworkConfiguration()
                                                    .withName("{vmss-name}")
                                                    .withPrimary(true)
                                                    .withIpConfigurations(
                                                        Arrays
                                                            .asList(
                                                                new VirtualMachineScaleSetIpConfiguration()
                                                                    .withName("{vmss-name}")
                                                                    .withSubnet(
                                                                        new ApiEntityReference()
                                                                            .withId(
                                                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"))))
                                                    .withEnableIpForwarding(true))))
                            .withApplicationProfile(
                                new ApplicationProfile()
                                    .withGalleryApplications(
                                        Arrays
                                            .asList(
                                                new VMGalleryApplication()
                                                    .withTags("myTag1")
                                                    .withOrder(1)
                                                    .withPackageReferenceId(
                                                        "/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0")
                                                    .withConfigurationReference(
                                                        "https://mystorageaccount.blob.core.windows.net/configurations/settings.config")
                                                    .withTreatFailureAsDeploymentFailure(true)
                                                    .withEnableAutomaticUpgrade(false),
                                                new VMGalleryApplication()
                                                    .withPackageReferenceId(
                                                        "/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1")))))
                    .withOverprovision(true),
                com.azure.core.util.Context.NONE);
    }
}
