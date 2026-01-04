
/**
 * Samples for WebApps GetAuthSettingsV2.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListAuthSettingsV2.json
     */
    /**
     * Sample code: List Auth Settings V2.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAuthSettingsV2(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getAuthSettingsV2WithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
