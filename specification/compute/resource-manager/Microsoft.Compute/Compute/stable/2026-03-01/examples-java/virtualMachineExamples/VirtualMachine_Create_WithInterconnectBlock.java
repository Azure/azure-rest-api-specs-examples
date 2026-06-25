
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.InterconnectBlockProfile;
import com.azure.resourcemanager.compute.models.InterconnectGroupProfile;
import com.azure.resourcemanager.compute.models.LinuxConfiguration;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithInterconnectBlock.json
     */
    /**
     * Sample code: Create or update a VM with Interconnect Block.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createOrUpdateAVMWithInterconnectBlock(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().createOrUpdate("myResourceGroup", "myVM",
            new VirtualMachineInner().withLocation("westus").withZones(Arrays.asList("1"))
                .withHardwareProfile(
                    new HardwareProfile().withVmSize(VirtualMachineSizeTypes.fromString("Standard_ND128isr_GB300_v6")))
                .withStorageProfile(new StorageProfile()
                    .withImageReference(new ImageReference().withPublisher("microsoft-dsvm").withOffer("ubuntu-hpc")
                        .withSku("2404-gb").withVersion("latest"))
                    .withOsDisk(new OSDisk().withName("myVMosdisk").withCaching(CachingTypes.READ_WRITE)
                        .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE).withManagedDisk(
                            new ManagedDiskParameters().withStorageAccountType(StorageAccountTypes.PREMIUM_LRS))))
                .withOsProfile(new OSProfile()
                    .withComputerName("myVM").withAdminUsername("{your-username}")
                    .withAdminPassword("fakeTokenPlaceholder")
                    .withLinuxConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(false)))
                .withNetworkProfile(new NetworkProfile().withNetworkInterfaces(
                    Arrays.asList(new NetworkInterfaceReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                        .withPrimary(true)))
                    .withInterconnectGroupProfile(new InterconnectGroupProfile().withInterconnectGroup(
                        new SubResource().withId(
                            "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup"))
                        .withSubgroups(Arrays.asList(new SubResource().withId(
                            "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0")))))
                .withInterconnectBlockProfile(
                    new InterconnectBlockProfile().withInterconnectBlock(new ApiEntityReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock"))),
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
