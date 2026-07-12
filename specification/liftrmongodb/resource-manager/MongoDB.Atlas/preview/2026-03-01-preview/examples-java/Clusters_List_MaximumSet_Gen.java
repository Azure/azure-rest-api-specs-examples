
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Clusters_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Clusters_List_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void clustersListMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.clusters().list("rgopenapi", "myOrganization", "myProject", com.azure.core.util.Context.NONE);
    }
}
