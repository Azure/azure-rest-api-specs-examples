
/**
 * Samples for PrivateLinks ListByMongoCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_PrivateLinkResourceList.json
     */
    /**
     * Sample code: Lists the private link resources available on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listsThePrivateLinkResourcesAvailableOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.privateLinks().listByMongoCluster("TestGroup", "myMongoCluster", com.azure.core.util.Context.NONE);
    }
}
