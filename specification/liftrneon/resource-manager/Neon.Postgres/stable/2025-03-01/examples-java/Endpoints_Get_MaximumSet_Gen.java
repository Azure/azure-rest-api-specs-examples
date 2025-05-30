
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Endpoints_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Endpoints_Get_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void endpointsGetMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.endpoints().getWithResponse("rgneon", "test-org", "entity-name", "entity-name", "entity-name",
            com.azure.core.util.Context.NONE);
    }
}
