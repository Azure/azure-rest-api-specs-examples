Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetParameters;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.SecurityEncryptionTypes;
import com.azure.resourcemanager.compute.models.SecurityProfile;
import com.azure.resourcemanager.compute.models.SecurityTypes;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.UefiSettings;
import com.azure.resourcemanager.compute.models.VMDiskSecurityProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;

/** Samples for VirtualMachines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAVmWithSecurityTypeConfidentialVMWithCustomerManagedKeys.json
     */
    /**
     * Sample code: Create a VM with securityType ConfidentialVM with Customer Managed Keys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVMWithSecurityTypeConfidentialVMWithCustomerManagedKeys(
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
                    .withHardwareProfile(
                        new HardwareProfile().withVmSize(VirtualMachineSizeTypes.fromString("Standard_DC2as_v5")))
                    .withStorageProfile(
                        new StorageProfile()
                            .withImageReference(
                                new ImageReference()
                                    .withPublisher("MicrosoftWindowsServer")
                                    .withOffer("2019-datacenter-cvm")
                                    .withSku("windows-cvm")
                                    .withVersion("17763.2183.2109130127"))
                            .withOsDisk(
                                new OSDisk()
                                    .withName("myVMosdisk")
                                    .withCaching(CachingTypes.READ_ONLY)
                                    .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                    .withManagedDisk(
                                        new ManagedDiskParameters()
                                            .withStorageAccountType(StorageAccountTypes.STANDARD_SSD_LRS)
                                            .withSecurityProfile(
                                                new VMDiskSecurityProfile()
                                                    .withSecurityEncryptionType(
                                                        SecurityEncryptionTypes.DISK_WITH_VMGUEST_STATE)
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
                                            .withPrimary(true))))
                    .withSecurityProfile(
                        new SecurityProfile()
                            .withUefiSettings(new UefiSettings().withSecureBootEnabled(true).withVTpmEnabled(true))
                            .withSecurityType(SecurityTypes.CONFIDENTIAL_VM)),
                Context.NONE);
    }
}
```
