Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DataDisk;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetParameters;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;

/** Samples for VirtualMachines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/CreateAVmWithDiskEncryptionSetResource.json
     */
    /**
     * Sample code: Create a vm with DiskEncryptionSet resource id in the os disk and data disk.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVmWithDiskEncryptionSetResourceIdInTheOsDiskAndDataDisk(
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
                                    .withId(
                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"))
                            .withOsDisk(
                                new OSDisk()
                                    .withName("myVMosdisk")
                                    .withCaching(CachingTypes.READ_WRITE)
                                    .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                    .withManagedDisk(
                                        new ManagedDiskParameters()
                                            .withStorageAccountType(StorageAccountTypes.STANDARD_LRS)
                                            .withDiskEncryptionSet(
                                                new DiskEncryptionSetParameters()
                                                    .withId(
                                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"))))
                            .withDataDisks(
                                Arrays
                                    .asList(
                                        new DataDisk()
                                            .withLun(0)
                                            .withCaching(CachingTypes.READ_WRITE)
                                            .withCreateOption(DiskCreateOptionTypes.EMPTY)
                                            .withDiskSizeGB(1023)
                                            .withManagedDisk(
                                                new ManagedDiskParameters()
                                                    .withStorageAccountType(StorageAccountTypes.STANDARD_LRS)
                                                    .withDiskEncryptionSet(
                                                        new DiskEncryptionSetParameters()
                                                            .withId(
                                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"))),
                                        new DataDisk()
                                            .withLun(1)
                                            .withCaching(CachingTypes.READ_WRITE)
                                            .withCreateOption(DiskCreateOptionTypes.ATTACH)
                                            .withDiskSizeGB(1023)
                                            .withManagedDisk(
                                                new ManagedDiskParameters()
                                                    .withId(
                                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-managed-disk-name}")
                                                    .withStorageAccountType(StorageAccountTypes.STANDARD_LRS)
                                                    .withDiskEncryptionSet(
                                                        new DiskEncryptionSetParameters()
                                                            .withId(
                                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"))))))
                    .withOsProfile(
                        new OSProfile()
                            .withComputerName("myVM")
                            .withAdminUsername("{your-username}")
                            .withAdminPassword("{your-password}"))
                    .withNetworkProfile(
                        new NetworkProfile()
                            .withNetworkInterfaces(
                                Arrays
                                    .asList(
                                        new NetworkInterfaceReference()
                                            .withId(
                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                                            .withPrimary(true)))),
                Context.NONE);
    }
}
```
