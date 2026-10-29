
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void operationsList(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
