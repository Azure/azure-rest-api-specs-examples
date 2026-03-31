
/**
 * Samples for WebApps GetBackupStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWebAppBackup.json
     */
    /**
     * Sample code: Get web app backup.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getWebAppBackup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getBackupStatusWithResponse("testrg123", "sitef6141", "12345",
            com.azure.core.util.Context.NONE);
    }
}
