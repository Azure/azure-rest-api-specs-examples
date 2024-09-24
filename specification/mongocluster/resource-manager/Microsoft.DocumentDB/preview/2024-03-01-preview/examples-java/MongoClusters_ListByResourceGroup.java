
/**
 * Samples for MongoClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists the Mongo Cluster resources in a resource group.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listsTheMongoClusterResourcesInAResourceGroup(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().listByResourceGroup("TestResourceGroup", com.azure.core.util.Context.NONE);
    }
}
