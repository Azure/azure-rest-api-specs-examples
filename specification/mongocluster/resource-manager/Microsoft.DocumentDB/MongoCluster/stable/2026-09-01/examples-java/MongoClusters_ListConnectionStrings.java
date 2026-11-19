
/**
 * Samples for MongoClusters ListConnectionStrings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_ListConnectionStrings.json
     */
    /**
     * Sample code: List the available connection strings for the DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listTheAvailableConnectionStringsForTheDocumentDBClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().listConnectionStringsWithResponse("TestGroup", "myMongoCluster",
            com.azure.core.util.Context.NONE);
    }
}
