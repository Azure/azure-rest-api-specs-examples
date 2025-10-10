
import com.azure.resourcemanager.mongocluster.models.CreateMode;
import com.azure.resourcemanager.mongocluster.models.MongoClusterProperties;
import com.azure.resourcemanager.mongocluster.models.MongoClusterReplicaParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/MongoClusters_CreateGeoReplica.json
     */
    /**
     * Sample code: Creates a replica Mongo Cluster resource from a source resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void createsAReplicaMongoClusterResourceFromASourceResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().define("myReplicaMongoCluster").withRegion("centralus")
            .withExistingResourceGroup("TestResourceGroup")
            .withProperties(new MongoClusterProperties().withCreateMode(CreateMode.GEO_REPLICA)
                .withReplicaParameters(new MongoClusterReplicaParameters().withSourceResourceId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/mySourceMongoCluster")
                    .withSourceLocation("eastus")))
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
