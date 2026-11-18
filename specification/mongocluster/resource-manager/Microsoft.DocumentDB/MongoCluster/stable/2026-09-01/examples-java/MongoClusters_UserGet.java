
/**
 * Samples for Users Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_UserGet.json
     */
    /**
     * Sample code: Gets a user on a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        getsAUserOnADocumentDBClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.users().getWithResponse("TestGroup", "myMongoCluster", "uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu",
            com.azure.core.util.Context.NONE);
    }
}
