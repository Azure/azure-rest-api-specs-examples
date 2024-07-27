
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DataDisk;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
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
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_Create_PlatformImageVmWithUnmanagedOsAndDataDisks.json
     */
    /**
     * Sample code: Create a platform-image vm with unmanaged os and data disks.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAPlatformImageVmWithUnmanagedOsAndDataDisks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines()
            .createOrUpdate("myResourceGroup", "{vm-name}", new VirtualMachineInner().withLocation("westus")
                .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D2_V2))
                .withStorageProfile(new StorageProfile()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2016-Datacenter").withVersion("latest"))
                    .withOsDisk(new OSDisk().withName("myVMosdisk").withVhd(new VirtualHardDisk().withUri(
                        "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"))
                        .withCaching(CachingTypes.READ_WRITE).withCreateOption(DiskCreateOptionTypes.FROM_IMAGE))
                    .withDataDisks(Arrays.asList(new DataDisk().withLun(0).withVhd(new VirtualHardDisk().withUri(
                        "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd"))
                        .withCreateOption(DiskCreateOptionTypes.EMPTY).withDiskSizeGB(1023),
                        new DataDisk().withLun(1).withVhd(new VirtualHardDisk().withUri(
                            "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd"))
                            .withCreateOption(DiskCreateOptionTypes.EMPTY).withDiskSizeGB(1023))))
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
