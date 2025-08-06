
import com.azure.resourcemanager.mongocluster.models.AdministratorProperties;
import com.azure.resourcemanager.mongocluster.models.AuthConfigProperties;
import com.azure.resourcemanager.mongocluster.models.AuthenticationMode;
import com.azure.resourcemanager.mongocluster.models.ComputeProperties;
import com.azure.resourcemanager.mongocluster.models.DataApiMode;
import com.azure.resourcemanager.mongocluster.models.DataApiProperties;
import com.azure.resourcemanager.mongocluster.models.HighAvailabilityMode;
import com.azure.resourcemanager.mongocluster.models.HighAvailabilityProperties;
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import com.azure.resourcemanager.mongocluster.models.PublicNetworkAccess;
import com.azure.resourcemanager.mongocluster.models.ShardingProperties;
import com.azure.resourcemanager.mongocluster.models.StorageProperties;
import com.azure.resourcemanager.mongocluster.models.StorageType;
import java.util.Arrays;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_Update.json
     */
    /**
     * Sample code: Updates a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        updatesAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new MongoClusterUpdateProperties()
            .withAdministrator(new AdministratorProperties().withUserName("mongoAdmin")).withServerVersion("5.0")
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withHighAvailability(new HighAvailabilityProperties().withTargetMode(HighAvailabilityMode.SAME_ZONE))
            .withStorage(new StorageProperties().withSizeGb(256L).withType(StorageType.PREMIUM_SSD))
            .withSharding(new ShardingProperties().withShardCount(4))
            .withCompute(new ComputeProperties().withTier("M50"))
            .withDataApi(new DataApiProperties().withMode(DataApiMode.DISABLED)).withPreviewFeatures(Arrays.asList())
            .withAuthConfig(new AuthConfigProperties().withAllowedModes(Arrays.asList(AuthenticationMode.NATIVE_AUTH))))
            .apply();
    }
}
