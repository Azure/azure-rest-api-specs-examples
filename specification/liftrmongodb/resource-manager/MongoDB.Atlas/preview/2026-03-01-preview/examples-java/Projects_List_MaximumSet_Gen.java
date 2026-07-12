
/**
 * Samples for Projects List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Projects_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_List_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void projectsListMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.projects().list("rgopenapi", "myOrganization", com.azure.core.util.Context.NONE);
    }
}
