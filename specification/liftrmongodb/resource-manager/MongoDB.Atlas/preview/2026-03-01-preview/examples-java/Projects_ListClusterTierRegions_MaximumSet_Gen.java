
/**
 * Samples for Projects ListClusterTierRegions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Projects_ListClusterTierRegions_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_ListClusterTierRegions_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        projectsListClusterTierRegionsMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.projects().listClusterTierRegionsWithResponse("rgopenapi", "myOrganization", "myProject",
            com.azure.core.util.Context.NONE);
    }
}
