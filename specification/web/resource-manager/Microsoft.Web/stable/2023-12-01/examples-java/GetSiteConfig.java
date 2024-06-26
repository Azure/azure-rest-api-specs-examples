
/**
 * Samples for WebApps GetConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetSiteConfig.json
     */
    /**
     * Sample code: Get Site Config.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSiteConfig(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getConfigurationWithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
