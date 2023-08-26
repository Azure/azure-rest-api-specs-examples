import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.LinuxConfiguration;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.SshConfiguration;
import com.azure.resourcemanager.compute.models.SshPublicKey;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
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

/** Samples for VirtualMachineScaleSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSshAuthentication.json
     */
    /**
     * Sample code: Create a scale set with ssh authentication.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAScaleSetWithSshAuthentication(com.azure.resourcemanager.AzureResourceManager azure) {
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
                                    .withLinuxConfiguration(
                                        new LinuxConfiguration()
                                            .withDisablePasswordAuthentication(true)
                                            .withSsh(
                                                new SshConfiguration()
                                                    .withPublicKeys(
                                                        Arrays
                                                            .asList(
                                                                new SshPublicKey()
                                                                    .withPath(
                                                                        "/home/{your-username}/.ssh/authorized_keys")
                                                                    .withKeyData("fakeTokenPlaceholder"))))))
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
                                                    .withEnableIpForwarding(true)))))
                    .withOverprovision(true),
                com.azure.core.util.Context.NONE);
    }
}
