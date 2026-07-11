
/**
 * Samples for Projects TierLimitReached.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Projects_TierLimitReached_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_TierLimitReached_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        projectsTierLimitReachedMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.projects().tierLimitReachedWithResponse("rgopenapi", "myOrganization", "myProject",
            com.azure.core.util.Context.NONE);
    }
}
