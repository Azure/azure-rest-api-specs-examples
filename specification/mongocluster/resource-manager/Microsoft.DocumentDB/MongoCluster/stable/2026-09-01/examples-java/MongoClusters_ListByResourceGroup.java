
/**
 * Samples for MongoClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists the DocumentDB cluster resources in a resource group.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listsTheDocumentDBClusterResourcesInAResourceGroup(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().listByResourceGroup("TestResourceGroup", com.azure.core.util.Context.NONE);
    }
}
