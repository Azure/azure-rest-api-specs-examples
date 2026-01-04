
/**
 * Samples for Global GetDeletedWebApp.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetDeletedWebApp.json
     */
    /**
     * Sample code: Get Deleted Web App.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDeletedWebApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getGlobals().getDeletedWebAppWithResponse("9",
            com.azure.core.util.Context.NONE);
    }
}
