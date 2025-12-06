
/**
 * Samples for ActiveDirectoryConfigs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ActiveDirectoryConfigs_Get.json
     */
    /**
     * Sample code: ActiveDirectoryConfigs_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void activeDirectoryConfigsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.activeDirectoryConfigs().getByResourceGroupWithResponse("myRG", "adconfig1",
            com.azure.core.util.Context.NONE);
    }
}
