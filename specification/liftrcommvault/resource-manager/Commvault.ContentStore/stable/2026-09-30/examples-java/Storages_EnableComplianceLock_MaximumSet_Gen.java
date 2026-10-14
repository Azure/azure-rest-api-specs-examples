
/**
 * Samples for Storages EnableComplianceLock.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-30/Storages_EnableComplianceLock_MaximumSet_Gen.json
     */
    /**
     * Sample code: Storages_EnableComplianceLock_MaximumSet - Enable compliance lock on storage.
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void storagesEnableComplianceLockMaximumSetEnableComplianceLockOnStorage(
        com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.storages().enableComplianceLockWithResponse("rgcommvault", "myCloudAccount", "myStorage",
            com.azure.core.util.Context.NONE);
    }
}
