import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DeleteOptions;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkApiVersion;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.PublicIpAddressSku;
import com.azure.resourcemanager.compute.models.PublicIpAddressSkuName;
import com.azure.resourcemanager.compute.models.PublicIpAddressSkuTier;
import com.azure.resourcemanager.compute.models.PublicIpAllocationMethod;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineNetworkInterfaceConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineNetworkInterfaceIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachinePublicIpAddressConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;

/** Samples for VirtualMachines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExamples/VirtualMachine_Create_WithNetworkInterfaceConfiguration.json
     */
    /**
     * Sample code: Create a VM with network interface configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVMWithNetworkInterfaceConfiguration(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .createOrUpdate(
                "myResourceGroup",
                "myVM",
                new VirtualMachineInner()
                    .withLocation("westus")
                    .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D1_V2))
                    .withStorageProfile(
                        new StorageProfile()
                            .withImageReference(
                                new ImageReference()
                                    .withPublisher("MicrosoftWindowsServer")
                                    .withOffer("WindowsServer")
                                    .withSku("2016-Datacenter")
                                    .withVersion("latest"))
                            .withOsDisk(
                                new OSDisk()
                                    .withName("myVMosdisk")
                                    .withCaching(CachingTypes.READ_WRITE)
                                    .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                    .withManagedDisk(
                                        new ManagedDiskParameters()
                                            .withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                    .withOsProfile(
                        new OSProfile()
                            .withComputerName("myVM")
                            .withAdminUsername("{your-username}")
                            .withAdminPassword("fakeTokenPlaceholder"))
                    .withNetworkProfile(
                        new NetworkProfile()
                            .withNetworkApiVersion(NetworkApiVersion.TWO_ZERO_TWO_ZERO_ONE_ONE_ZERO_ONE)
                            .withNetworkInterfaceConfigurations(
                                Arrays
                                    .asList(
                                        new VirtualMachineNetworkInterfaceConfiguration()
                                            .withName("{nic-config-name}")
                                            .withPrimary(true)
                                            .withDeleteOption(DeleteOptions.DELETE)
                                            .withIpConfigurations(
                                                Arrays
                                                    .asList(
                                                        new VirtualMachineNetworkInterfaceIpConfiguration()
                                                            .withName("{ip-config-name}")
                                                            .withPrimary(true)
                                                            .withPublicIpAddressConfiguration(
                                                                new VirtualMachinePublicIpAddressConfiguration()
                                                                    .withName("{publicIP-config-name}")
                                                                    .withSku(
                                                                        new PublicIpAddressSku()
                                                                            .withName(PublicIpAddressSkuName.BASIC)
                                                                            .withTier(PublicIpAddressSkuTier.GLOBAL))
                                                                    .withDeleteOption(DeleteOptions.DETACH)
                                                                    .withPublicIpAllocationMethod(
                                                                        PublicIpAllocationMethod.STATIC))))))),
                com.azure.core.util.Context.NONE);
    }
}
