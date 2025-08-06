
import com.azure.resourcemanager.mongocluster.models.AdministratorProperties;
import com.azure.resourcemanager.mongocluster.models.AuthConfigProperties;
import com.azure.resourcemanager.mongocluster.models.AuthenticationMode;
import com.azure.resourcemanager.mongocluster.models.ComputeProperties;
import com.azure.resourcemanager.mongocluster.models.HighAvailabilityMode;
import com.azure.resourcemanager.mongocluster.models.HighAvailabilityProperties;
import com.azure.resourcemanager.mongocluster.models.MongoClusterProperties;
import com.azure.resourcemanager.mongocluster.models.ShardingProperties;
import com.azure.resourcemanager.mongocluster.models.StorageProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_Create.json
     */
    /**
     * Sample code: Creates a new Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        createsANewMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().define("myMongoCluster").withRegion("westus2")
            .withExistingResourceGroup("TestResourceGroup")
            .withProperties(new MongoClusterProperties()
                .withAdministrator(
                    new AdministratorProperties().withUserName("mongoAdmin").withPassword("fakeTokenPlaceholder"))
                .withServerVersion("5.0")
                .withHighAvailability(
                    new HighAvailabilityProperties().withTargetMode(HighAvailabilityMode.ZONE_REDUNDANT_PREFERRED))
                .withStorage(new StorageProperties().withSizeGb(128L))
                .withSharding(new ShardingProperties().withShardCount(1))
                .withCompute(new ComputeProperties().withTier("M30")).withAuthConfig(
                    new AuthConfigProperties().withAllowedModes(Arrays.asList(AuthenticationMode.NATIVE_AUTH))))
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
