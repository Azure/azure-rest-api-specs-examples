
/**
 * Samples for Projects Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Projects_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_Delete_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void projectsDeleteMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.projects().deleteWithResponse("rgneon", "test-org", "entity-name", com.azure.core.util.Context.NONE);
    }
}
