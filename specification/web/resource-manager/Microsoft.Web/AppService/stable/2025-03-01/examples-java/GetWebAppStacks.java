
/**
 * Samples for Provider GetWebAppStacks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetWebAppStacks.json
     */
    /**
     * Sample code: Get Web App Stacks.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getWebAppStacks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getProviders().getWebAppStacks(null,
            com.azure.core.util.Context.NONE);
    }
}
