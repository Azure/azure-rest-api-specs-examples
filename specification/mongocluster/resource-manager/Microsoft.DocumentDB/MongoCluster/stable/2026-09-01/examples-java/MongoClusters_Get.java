
/**
 * Samples for MongoClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_Get.json
     */
    /**
     * Sample code: Gets a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        getsADocumentDBClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster",
            com.azure.core.util.Context.NONE);
    }
}
