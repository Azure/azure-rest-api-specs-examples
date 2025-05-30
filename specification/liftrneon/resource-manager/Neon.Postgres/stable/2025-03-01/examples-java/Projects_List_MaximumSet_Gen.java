
/**
 * Samples for Projects List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Projects_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_List_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void projectsListMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.projects().list("rgneon", "test-org", com.azure.core.util.Context.NONE);
    }
}
