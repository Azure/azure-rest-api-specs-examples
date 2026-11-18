
/**
 * Samples for MongoClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_List.json
     */
    /**
     * Sample code: Lists the DocumentDB cluster resources in a subscription.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listsTheDocumentDBClusterResourcesInASubscription(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().list(com.azure.core.util.Context.NONE);
    }
}
