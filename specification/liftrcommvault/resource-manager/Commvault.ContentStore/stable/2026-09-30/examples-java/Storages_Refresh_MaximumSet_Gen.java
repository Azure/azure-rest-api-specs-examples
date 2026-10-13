
/**
 * Samples for Storages Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-30/Storages_Refresh_MaximumSet_Gen.json
     */
    /**
     * Sample code: Storages_Refresh_MaximumSet - Refresh storage state from Commvault.
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void storagesRefreshMaximumSetRefreshStorageStateFromCommvault(
        com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.storages().refreshWithResponse("rgcommvault", "myCloudAccount", "myStorage",
            com.azure.core.util.Context.NONE);
    }
}
