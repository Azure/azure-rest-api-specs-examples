
/**
 * Samples for WebApps ListSiteBackupsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListSiteBackupsSlot.json
     */
    /**
     * Sample code: List backups.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listBackups(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listSiteBackupsSlot("testrg123", "tests346", "staging",
            com.azure.core.util.Context.NONE);
    }
}
