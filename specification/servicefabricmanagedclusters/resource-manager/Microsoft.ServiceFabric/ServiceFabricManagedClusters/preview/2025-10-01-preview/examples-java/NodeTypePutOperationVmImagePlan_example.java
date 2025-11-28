
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmImagePlan;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NodeTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationVmImagePlan_example.json
     */
    /**
     * Sample code: Put node type with vm image plan.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putNodeTypeWithVmImagePlan(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().define("BE").withExistingManagedCluster("resRg", "myCluster").withIsPrimary(false)
            .withVmInstanceCount(10).withDataDiskSizeGB(200).withVmSize("Standard_D3")
            .withVmImagePublisher("testpublisher").withVmImageOffer("windows_2022_test")
            .withVmImageSku("win_2022_test_20_10_gen2").withVmImageVersion("latest").withVmImagePlan(new VmImagePlan()
                .withName("win_2022_test_20_10_gen2").withProduct("windows_2022_test").withPublisher("testpublisher"))
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
