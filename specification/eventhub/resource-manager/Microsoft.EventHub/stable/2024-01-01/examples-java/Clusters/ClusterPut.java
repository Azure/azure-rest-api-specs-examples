
import com.azure.resourcemanager.eventhubs.fluent.models.ClusterInner;
import com.azure.resourcemanager.eventhubs.models.ClusterSku;
import com.azure.resourcemanager.eventhubs.models.ClusterSkuName;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/ClusterPut.json
     */
    /**
     * Sample code: ClusterPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void clusterPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().createOrUpdate("myResourceGroup", "testCluster",
            new ClusterInner().withLocation("South Central US").withTags(mapOf("tag1", "value1", "tag2", "value2"))
                .withSku(new ClusterSku().withName(ClusterSkuName.DEDICATED).withCapacity(1)),
            com.azure.core.util.Context.NONE);
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
