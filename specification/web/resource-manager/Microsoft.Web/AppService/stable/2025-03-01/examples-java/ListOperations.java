
/**
 * Samples for Provider ListOperations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListOperations.json
     */
    /**
     * Sample code: List operations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getProviders().listOperations(com.azure.core.util.Context.NONE);
    }
}
