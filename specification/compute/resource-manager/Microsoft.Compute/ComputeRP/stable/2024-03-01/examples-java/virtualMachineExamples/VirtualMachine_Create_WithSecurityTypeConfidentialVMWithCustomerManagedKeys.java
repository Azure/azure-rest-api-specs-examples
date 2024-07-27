
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
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_Create_WithSecurityTypeConfidentialVMWithCustomerManagedKeys.json
     */
    /**
     * Sample code: Create a VM with securityType ConfidentialVM with Customer Managed Keys.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVMWithSecurityTypeConfidentialVMWithCustomerManagedKeys(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().createOrUpdate("myResourceGroup", "myVM",
            new VirtualMachineInner().withLocation("westus")
                .withHardwareProfile(
                    new HardwareProfile().withVmSize(VirtualMachineSizeTypes.fromString("Standard_DC2as_v5")))
                .withStorageProfile(new StorageProfile()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("2019-datacenter-cvm").withSku("windows-cvm").withVersion("17763.2183.2109130127"))
                    .withOsDisk(new OSDisk().withName("myVMosdisk").withCaching(CachingTypes.READ_ONLY)
                        .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                        .withManagedDisk(new ManagedDiskParameters()
                            .withStorageAccountType(StorageAccountTypes.STANDARD_SSD_LRS)
                            .withSecurityProfile(new VMDiskSecurityProfile()
                                .withSecurityEncryptionType(SecurityEncryptionTypes.DISK_WITH_VMGUEST_STATE)
                                .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId(
                                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"))))))
                .withOsProfile(new OSProfile().withComputerName("myVM").withAdminUsername("{your-username}")
                    .withAdminPassword("fakeTokenPlaceholder"))
                .withNetworkProfile(
                    new NetworkProfile().withNetworkInterfaces(Arrays.asList(new NetworkInterfaceReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                        .withPrimary(true))))
                .withSecurityProfile(new SecurityProfile()
                    .withUefiSettings(new UefiSettings().withSecureBootEnabled(true).withVTpmEnabled(true))
                    .withSecurityType(SecurityTypes.CONFIDENTIAL_VM)),
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
