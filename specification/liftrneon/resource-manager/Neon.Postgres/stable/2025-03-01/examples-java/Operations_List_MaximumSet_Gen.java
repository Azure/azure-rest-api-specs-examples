
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void operationsListMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
