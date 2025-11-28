
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NodeTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationCustomImage_example.json
     */
    /**
     * Sample code: Put node type with custom vm image.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putNodeTypeWithCustomVmImage(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().define("BE").withExistingManagedCluster("resRg", "myCluster").withIsPrimary(false)
            .withVmInstanceCount(10).withDataDiskSizeGB(200).withVmSize("Standard_D3")
            .withVmImageResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-custom-image/providers/Microsoft.Compute/galleries/myCustomImages/images/Win2019DC")
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
