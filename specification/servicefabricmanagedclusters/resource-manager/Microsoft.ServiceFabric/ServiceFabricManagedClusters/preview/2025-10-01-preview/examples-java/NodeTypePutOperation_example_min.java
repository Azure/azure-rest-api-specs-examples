
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NodeTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperation_example_min.json
     */
    /**
     * Sample code: Put a node type with minimum parameters.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putANodeTypeWithMinimumParameters(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().define("BE").withExistingManagedCluster("resRg", "myCluster").withIsPrimary(false)
            .withVmInstanceCount(10).withDataDiskSizeGB(200).withVmSize("Standard_D3")
            .withVmImagePublisher("MicrosoftWindowsServer").withVmImageOffer("WindowsServer")
            .withVmImageSku("2016-Datacenter-Server-Core").withVmImageVersion("latest").create();
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
