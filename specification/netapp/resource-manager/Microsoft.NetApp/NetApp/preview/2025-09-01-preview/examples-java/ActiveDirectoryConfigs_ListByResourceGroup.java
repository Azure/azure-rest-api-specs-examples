
/**
 * Samples for ActiveDirectoryConfigs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ActiveDirectoryConfigs_ListByResourceGroup.json
     */
    /**
     * Sample code: ActiveDirectoryConfigs_ListByResourceGroup.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        activeDirectoryConfigsListByResourceGroup(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.activeDirectoryConfigs().listByResourceGroup("myRG", com.azure.core.util.Context.NONE);
    }
}
