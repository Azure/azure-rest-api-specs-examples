
/**
 * Samples for WebApps DeleteBackup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteWebAppBackup.json
     */
    /**
     * Sample code: Delete web app backup.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteWebAppBackup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().deleteBackupWithResponse("testrg123", "sitef6141", "12345",
            com.azure.core.util.Context.NONE);
    }
}
