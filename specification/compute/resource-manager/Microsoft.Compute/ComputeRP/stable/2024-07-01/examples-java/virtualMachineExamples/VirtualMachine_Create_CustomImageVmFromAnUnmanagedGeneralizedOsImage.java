
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualHardDisk;
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
     * virtualMachineExamples/VirtualMachine_Create_CustomImageVmFromAnUnmanagedGeneralizedOsImage.json
     */
    /**
     * Sample code: Create a custom-image vm from an unmanaged generalized os image.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createACustomImageVmFromAnUnmanagedGeneralizedOsImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines()
            .createOrUpdate("myResourceGroup", "{vm-name}", new VirtualMachineInner().withLocation("westus")
                .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D1_V2))
                .withStorageProfile(new StorageProfile().withOsDisk(new OSDisk()
                    .withOsType(OperatingSystemTypes.WINDOWS).withName("myVMosdisk")
                    .withVhd(new VirtualHardDisk().withUri(
                        "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"))
                    .withImage(new VirtualHardDisk().withUri(
                        "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd"))
                    .withCaching(CachingTypes.READ_WRITE).withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)))
                .withOsProfile(new OSProfile().withComputerName("myVM").withAdminUsername("{your-username}")
                    .withAdminPassword("fakeTokenPlaceholder"))
                .withNetworkProfile(
                    new NetworkProfile().withNetworkInterfaces(Arrays.asList(new NetworkInterfaceReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                        .withPrimary(true)))),
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
