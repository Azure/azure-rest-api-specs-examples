
/**
 * Samples for Branches Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Branches_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Branches_Get_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void branchesGetMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.branches().getWithResponse("rgneon", "test-org", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
