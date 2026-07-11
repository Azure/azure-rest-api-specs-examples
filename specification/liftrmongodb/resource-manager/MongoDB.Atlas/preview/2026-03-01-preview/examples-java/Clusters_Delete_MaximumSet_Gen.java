
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Clusters_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Clusters_Delete_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void clustersDeleteMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.clusters().delete("rgopenapi", "myOrganization", "myProject", "myCluster",
            com.azure.core.util.Context.NONE);
    }
}
