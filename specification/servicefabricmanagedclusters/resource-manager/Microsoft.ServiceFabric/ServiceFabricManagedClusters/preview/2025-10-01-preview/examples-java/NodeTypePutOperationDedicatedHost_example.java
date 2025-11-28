
import com.azure.resourcemanager.servicefabricmanagedclusters.models.DiskType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NodeTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationDedicatedHost_example.json
     */
    /**
     * Sample code: Put node type with dedicated hosts.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putNodeTypeWithDedicatedHosts(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().define("BE").withExistingManagedCluster("resRg", "myCluster").withIsPrimary(false)
            .withVmInstanceCount(10).withDataDiskSizeGB(200).withDataDiskType(DiskType.STANDARD_SSD_LRS)
            .withPlacementProperties(mapOf()).withCapacities(mapOf()).withVmSize("Standard_D8s_v3")
            .withVmImagePublisher("MicrosoftWindowsServer").withVmImageOffer("WindowsServer")
            .withVmImageSku("2019-Datacenter").withVmImageVersion("latest").withZones(Arrays.asList("1"))
            .withHostGroupId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testhostgroupRG/providers/Microsoft.Compute/hostGroups/testHostGroup")
            .create();
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
