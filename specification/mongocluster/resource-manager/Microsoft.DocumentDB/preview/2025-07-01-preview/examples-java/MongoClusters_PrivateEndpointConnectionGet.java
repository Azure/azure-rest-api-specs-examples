
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Get a private endpoint connection on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void getAPrivateEndpointConnectionOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.privateEndpointConnections().getWithResponse("TestGroup", "myMongoCluster",
            "pecTest.5d393f64-ef64-46d0-9959-308321c44ac0", com.azure.core.util.Context.NONE);
    }
}
