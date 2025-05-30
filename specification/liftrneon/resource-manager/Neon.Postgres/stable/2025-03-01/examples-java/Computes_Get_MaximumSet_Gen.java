
/**
 * Samples for Computes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Computes_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Computes_Get_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void computesGetMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.computes().getWithResponse("rgneon", "test-org", "entity-name", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
