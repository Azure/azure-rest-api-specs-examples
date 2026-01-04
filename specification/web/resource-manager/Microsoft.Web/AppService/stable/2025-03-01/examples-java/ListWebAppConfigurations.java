
/**
 * Samples for WebApps ListConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWebAppConfigurations.
     * json
     */
    /**
     * Sample code: List web app configurations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listWebAppConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listConfigurations("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
