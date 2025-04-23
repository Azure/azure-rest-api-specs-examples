
/**
 * Samples for Branches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Branches_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Branches_Delete_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void branchesDeleteMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.branches().deleteWithResponse("rgneon", "test-org", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
