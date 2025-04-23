
/**
 * Samples for NeonDatabases List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonDatabases_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonDatabases_List_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void neonDatabasesListMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.neonDatabases().list("rgneon", "test-org", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
