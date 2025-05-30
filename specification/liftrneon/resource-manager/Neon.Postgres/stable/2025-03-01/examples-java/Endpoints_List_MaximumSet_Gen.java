
/**
 * Samples for Endpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Endpoints_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Endpoints_List_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void endpointsListMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.endpoints().list("rgneon", "test-org", "entity-name", "entity-name", com.azure.core.util.Context.NONE);
    }
}
