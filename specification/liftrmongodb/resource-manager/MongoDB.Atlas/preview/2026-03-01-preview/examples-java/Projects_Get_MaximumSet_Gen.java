
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Projects_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_Get_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void projectsGetMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.projects().getWithResponse("rgopenapi", "myOrganization", "myProject",
            com.azure.core.util.Context.NONE);
    }
}
