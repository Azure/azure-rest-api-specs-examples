
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.AdditionalDiskProperties;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DataDisk;
import com.azure.resourcemanager.compute.models.DiskApiVersion;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineDiskNetworkAccessPolicy;
import com.azure.resourcemanager.compute.models.VirtualMachineDiskProperties;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Create_WithAdditionalDiskProperties.json
     */
    /**
     * Sample code: Create a vm with additional disk properties (network access policy and tier).
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createAVmWithAdditionalDiskPropertiesNetworkAccessPolicyAndTier(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().createOrUpdate("myResourceGroup", "myVM", new VirtualMachineInner()
            .withLocation("westus")
            .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D4S_V3))
            .withStorageProfile(new StorageProfile()
                .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                    .withOffer("WindowsServer").withSku("2022-datacenter-azure-edition").withVersion("latest"))
                .withOsDisk(new OSDisk().withName("myVMosdisk").withCaching(CachingTypes.READ_WRITE)
                    .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                    .withManagedDisk(new ManagedDiskParameters().withStorageAccountType(StorageAccountTypes.PREMIUM_LRS)
                        .withAdditionalDiskProperties(
                            new AdditionalDiskProperties().withManagedDiskProperties(new VirtualMachineDiskProperties()
                                .withNetworkAccessPolicy(VirtualMachineDiskNetworkAccessPolicy.ALLOW_PRIVATE)
                                .withDiskAccessId(
                                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess")))))
                .withDataDisks(Arrays.asList(new DataDisk().withLun(0).withName("myDataDisk")
                    .withCreateOption(DiskCreateOptionTypes.EMPTY).withDiskSizeGB(1024)
                    .withManagedDisk(new ManagedDiskParameters().withStorageAccountType(StorageAccountTypes.PREMIUM_LRS)
                        .withAdditionalDiskProperties(new AdditionalDiskProperties()
                            .withManagedDiskProperties(new VirtualMachineDiskProperties().withTier("P30")
                                .withNetworkAccessPolicy(VirtualMachineDiskNetworkAccessPolicy.ALLOW_PRIVATE)
                                .withDiskAccessId(
                                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess"))))))
                .withDiskApiVersion(DiskApiVersion.TWO_ZERO_TWO_SIX_ZERO_THREE_ZERO_TWO))
            .withOsProfile(new OSProfile().withComputerName("myVM").withAdminUsername("{your-username}")
                .withAdminPassword("fakeTokenPlaceholder"))
            .withNetworkProfile(new NetworkProfile().withNetworkInterfaces(Arrays.asList(new NetworkInterfaceReference()
                .withId(
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
