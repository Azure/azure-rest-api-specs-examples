
/**
 * Samples for Computes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Computes_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Computes_Delete_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void computesDeleteMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.computes().deleteWithResponse("rgneon", "test-org", "entity-name", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
