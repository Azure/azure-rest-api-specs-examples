
/**
 * Samples for NeonRoles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonRoles_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonRoles_Delete_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void neonRolesDeleteMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.neonRoles().deleteWithResponse("rgneon", "test-org", "entity-name", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
