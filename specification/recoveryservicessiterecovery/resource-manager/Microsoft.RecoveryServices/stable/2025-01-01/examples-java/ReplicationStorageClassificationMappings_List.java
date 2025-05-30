
/**
 * Samples for StorageClassificationMappings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationStorageClassificationMappings_List.json
     */
    /**
     * Sample code: Gets the list of storage classification mappings objects under a vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfStorageClassificationMappingsObjectsUnderAVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.storageClassificationMappings().list("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
