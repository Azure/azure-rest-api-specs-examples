
/**
 * Samples for WebApps DeleteBackup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/DeleteWebAppBackup.json
     */
    /**
     * Sample code: Delete web app backup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteWebAppBackup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().deleteBackupWithResponse("testrg123", "sitef6141",
            "12345", com.azure.core.util.Context.NONE);
    }
}
