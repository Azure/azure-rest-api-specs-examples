
/**
 * Samples for PrivateEndpointConnections ListByMongoCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Lists the private endpoint connection resources on a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listsThePrivateEndpointConnectionResourcesOnADocumentDBClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.privateEndpointConnections().listByMongoCluster("TestGroup", "myMongoCluster",
            com.azure.core.util.Context.NONE);
    }
}
