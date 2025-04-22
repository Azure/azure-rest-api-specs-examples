
/**
 * Samples for Branches List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Branches_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Branches_List_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void branchesListMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.branches().list("rgneon", "test-org", "entity-name", com.azure.core.util.Context.NONE);
    }
}
