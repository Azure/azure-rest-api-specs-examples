
/**
 * Samples for Clusters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Clusters_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Clusters_Get_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void clustersGetMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.clusters().getWithResponse("rgopenapi", "myOrganization", "myProject", "myCluster",
            com.azure.core.util.Context.NONE);
    }
}
