
/**
 * Samples for Projects Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Projects_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_Delete_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void projectsDeleteMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.projects().delete("rgopenapi", "myOrganization", "myProject", com.azure.core.util.Context.NONE);
    }
}
