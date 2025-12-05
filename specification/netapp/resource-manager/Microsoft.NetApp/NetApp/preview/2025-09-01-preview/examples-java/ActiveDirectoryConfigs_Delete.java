
/**
 * Samples for ActiveDirectoryConfigs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ActiveDirectoryConfigs_Delete.json
     */
    /**
     * Sample code: ActiveDirectoryConfigs_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void activeDirectoryConfigsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.activeDirectoryConfigs().delete("myRG", "adconfig1", com.azure.core.util.Context.NONE);
    }
}
