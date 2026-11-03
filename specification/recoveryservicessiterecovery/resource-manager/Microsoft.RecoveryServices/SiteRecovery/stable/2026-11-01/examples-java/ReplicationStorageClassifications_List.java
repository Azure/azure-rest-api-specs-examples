
/**
 * Samples for StorageClassifications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationStorageClassifications_List.json
     */
    /**
     * Sample code: Gets the list of storage classification objects under a vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfStorageClassificationObjectsUnderAVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.storageClassifications().list("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
