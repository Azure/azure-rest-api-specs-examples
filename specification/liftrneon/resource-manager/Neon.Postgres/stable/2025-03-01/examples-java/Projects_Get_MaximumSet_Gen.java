
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Projects_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_Get_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void projectsGetMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.projects().getWithResponse("rgneon", "test-org", "entity-name", com.azure.core.util.Context.NONE);
    }
}
