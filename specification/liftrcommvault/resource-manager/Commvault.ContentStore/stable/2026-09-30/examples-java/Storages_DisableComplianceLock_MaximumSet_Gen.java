
/**
 * Samples for Storages DisableComplianceLock.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-30/Storages_DisableComplianceLock_MaximumSet_Gen.json
     */
    /**
     * Sample code: Storages_DisableComplianceLock_MaximumSet - Disable compliance lock on storage (initiates MPA
     * approval; status becomes DisablementPending).
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void
        storagesDisableComplianceLockMaximumSetDisableComplianceLockOnStorageInitiatesMPAApprovalStatusBecomesDisablementPending(
            com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.storages().disableComplianceLockWithResponse("rgcommvault", "myCloudAccount", "myStorage",
            com.azure.core.util.Context.NONE);
    }
}
