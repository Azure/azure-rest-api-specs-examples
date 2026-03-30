
/**
 * Samples for Provider ListOperations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListOperations.json
     */
    /**
     * Sample code: List operations.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listOperations(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getProviders().listOperations(com.azure.core.util.Context.NONE);
    }
}
