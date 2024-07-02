
/**
 * Samples for MongoClusters ListConnectionStrings.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mongocluster/DocumentDB.MongoCluster.Management/examples/2024-03-01-preview/
     * MongoClusters_ListConnectionStrings.json
     */
    /**
     * Sample code: List the available connection strings for the Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listTheAvailableConnectionStringsForTheMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().listConnectionStringsWithResponse("TestGroup", "myMongoCluster",
            com.azure.core.util.Context.NONE);
    }
}
