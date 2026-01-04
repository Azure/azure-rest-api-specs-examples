
/**
 * Samples for WebApps GetAuthSettings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListAuthSettings.json
     */
    /**
     * Sample code: List Auth Settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAuthSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getAuthSettingsWithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
