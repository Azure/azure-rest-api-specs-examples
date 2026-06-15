
/**
 * Samples for Disks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Disks_List.json
     */
    /**
     * Sample code: Disks_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void disksList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.disks().list("resourceGroupName", "{labName}", "@me", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
