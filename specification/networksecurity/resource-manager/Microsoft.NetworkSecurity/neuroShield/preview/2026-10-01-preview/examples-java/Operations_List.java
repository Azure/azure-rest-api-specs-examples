
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void operationsList(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
