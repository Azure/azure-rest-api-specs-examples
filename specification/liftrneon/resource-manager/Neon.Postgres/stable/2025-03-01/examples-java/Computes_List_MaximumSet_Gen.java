
/**
 * Samples for Computes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Computes_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Computes_List_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void computesListMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.computes().list("rgneon", "test-org", "entity-name", "entity-name", com.azure.core.util.Context.NONE);
    }
}
