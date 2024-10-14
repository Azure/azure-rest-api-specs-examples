
import com.azure.resourcemanager.mongocluster.models.CreateMode;
import com.azure.resourcemanager.mongocluster.models.MongoClusterProperties;
import com.azure.resourcemanager.mongocluster.models.MongoClusterReplicaParameters;

/**
 * Samples for MongoClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_CreateGeoReplica.json
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
}
