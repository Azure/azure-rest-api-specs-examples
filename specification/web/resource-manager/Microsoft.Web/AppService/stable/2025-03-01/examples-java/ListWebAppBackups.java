
/**
 * Samples for WebApps ListBackups.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWebAppBackups.json
     */
    /**
     * Sample code: List web app backups.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listWebAppBackups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listBackups("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
