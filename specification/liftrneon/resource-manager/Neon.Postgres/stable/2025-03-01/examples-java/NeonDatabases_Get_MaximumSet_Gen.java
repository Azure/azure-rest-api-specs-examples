
/**
 * Samples for NeonDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonDatabases_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonDatabases_Get_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void neonDatabasesGetMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.neonDatabases().getWithResponse("rgneon", "test-org", "entity-name", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
