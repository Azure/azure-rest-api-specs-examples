
/**
 * Samples for ActiveDirectoryConfigs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ActiveDirectoryConfigs_ListBySubscription.json
     */
    /**
     * Sample code: ActiveDirectoryConfigs_ListBySubscription.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        activeDirectoryConfigsListBySubscription(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.activeDirectoryConfigs().list(com.azure.core.util.Context.NONE);
    }
}
