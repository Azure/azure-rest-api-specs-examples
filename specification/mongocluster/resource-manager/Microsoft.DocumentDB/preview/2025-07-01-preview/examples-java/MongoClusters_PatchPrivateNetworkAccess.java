
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import com.azure.resourcemanager.mongocluster.models.PublicNetworkAccess;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_PatchPrivateNetworkAccess.json
     */
    /**
     * Sample code: Disables public network access on a Mongo Cluster resource with a private endpoint connection.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void disablesPublicNetworkAccessOnAMongoClusterResourceWithAPrivateEndpointConnection(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(new MongoClusterUpdateProperties().withPublicNetworkAccess(PublicNetworkAccess.DISABLED))
            .apply();
    }
}
