
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Delete a private endpoint connection on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void deleteAPrivateEndpointConnectionOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.privateEndpointConnections().delete("TestGroup", "myMongoCluster",
            "pecTest.5d393f64-ef64-46d0-9959-308321c44ac0", com.azure.core.util.Context.NONE);
    }
}
