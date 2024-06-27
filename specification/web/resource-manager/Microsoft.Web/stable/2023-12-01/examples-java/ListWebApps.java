
/**
 * Samples for WebApps List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListWebApps.json
     */
    /**
     * Sample code: List Web apps for subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listWebAppsForSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().list(com.azure.core.util.Context.NONE);
    }
}
