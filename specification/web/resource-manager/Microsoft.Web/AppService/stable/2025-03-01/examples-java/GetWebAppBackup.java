
/**
 * Samples for WebApps GetBackupStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetWebAppBackup.json
     */
    /**
     * Sample code: Get web app backup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getWebAppBackup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getBackupStatusWithResponse("testrg123", "sitef6141",
            "12345", com.azure.core.util.Context.NONE);
    }
}
