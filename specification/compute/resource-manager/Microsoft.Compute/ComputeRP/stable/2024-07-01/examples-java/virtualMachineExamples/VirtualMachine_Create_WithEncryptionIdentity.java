
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.EncryptionIdentity;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.ResourceIdentityType;
import com.azure.resourcemanager.compute.models.SecurityProfile;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineIdentity;
import com.azure.resourcemanager.compute.models.VirtualMachineIdentityUserAssignedIdentities;
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
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineExamples/VirtualMachine_Create_WithEncryptionIdentity.json
     */
    /**
     * Sample code: Create a VM with encryption identity.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVMWithEncryptionIdentity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().createOrUpdate("myResourceGroup", "myVM",
            new VirtualMachineInner().withLocation("westus").withIdentity(new VirtualMachineIdentity()
                .withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
                    new VirtualMachineIdentityUserAssignedIdentities())))
                .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D2S_V3))
                .withStorageProfile(new StorageProfile()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2019-Datacenter").withVersion("latest"))
                    .withOsDisk(new OSDisk().withName("myVMosdisk").withCaching(CachingTypes.READ_ONLY)
                        .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE).withManagedDisk(
                            new ManagedDiskParameters().withStorageAccountType(StorageAccountTypes.STANDARD_SSD_LRS))))
                .withOsProfile(new OSProfile().withComputerName("myVM").withAdminUsername("{your-username}")
                    .withAdminPassword("fakeTokenPlaceholder"))
                .withNetworkProfile(
                    new NetworkProfile().withNetworkInterfaces(Arrays.asList(new NetworkInterfaceReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                        .withPrimary(true))))
                .withSecurityProfile(new SecurityProfile()
                    .withEncryptionIdentity(new EncryptionIdentity().withUserAssignedIdentityResourceId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"))),
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
