
/**
 * Samples for WebApps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/DeleteWebApp.json
     */
    /**
     * Sample code: Delete Web app.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteWebApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().deleteWithResponse("testrg123", "sitef6141", null, null,
            com.azure.core.util.Context.NONE);
    }
}
