
/**
 * Samples for Endpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Endpoints_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Endpoints_Delete_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void endpointsDeleteMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.endpoints().deleteWithResponse("rgneon", "test-org", "entity-name", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
