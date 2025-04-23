
/**
 * Samples for Organizations GetPostgresVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Organizations_GetPostgresVersions_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organizations_GetPostgresVersions_MinimumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        organizationsGetPostgresVersionsMinimumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().getPostgresVersionsWithResponse("rgneon", null, com.azure.core.util.Context.NONE);
    }
}
