
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-01-preview/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void operationsList(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
