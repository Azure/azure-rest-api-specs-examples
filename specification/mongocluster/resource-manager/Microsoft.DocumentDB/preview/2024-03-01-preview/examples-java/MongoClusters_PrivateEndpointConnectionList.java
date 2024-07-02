
/**
 * Samples for PrivateEndpointConnections ListByMongoCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mongocluster/DocumentDB.MongoCluster.Management/examples/2024-03-01-preview/
     * MongoClusters_PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Lists the private endpoint connection resources on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listsThePrivateEndpointConnectionResourcesOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.privateEndpointConnections().listByMongoCluster("TestGroup", "myMongoCluster",
            com.azure.core.util.Context.NONE);
    }
}
