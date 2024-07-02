
/**
 * Samples for MongoClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mongocluster/DocumentDB.MongoCluster.Management/examples/2024-03-01-preview/MongoClusters_List.json
     */
    /**
     * Sample code: Lists the Mongo Cluster resources in a subscription.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listsTheMongoClusterResourcesInASubscription(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().list(com.azure.core.util.Context.NONE);
    }
}
