
/**
 * Samples for WebApps ListApplicationSettings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListAppSettings.json
     */
    /**
     * Sample code: List App Settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listApplicationSettingsWithResponse("testrg123",
            "sitef6141", com.azure.core.util.Context.NONE);
    }
}
