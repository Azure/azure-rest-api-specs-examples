
/**
 * Samples for NeonRoles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonRoles_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonRoles_List_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void neonRolesListMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.neonRoles().list("rgneon", "test-org", "entity-name", "entity-name", com.azure.core.util.Context.NONE);
    }
}
