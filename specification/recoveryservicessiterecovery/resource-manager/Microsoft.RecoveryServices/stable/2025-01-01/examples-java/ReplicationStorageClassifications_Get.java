
/**
 * Samples for StorageClassifications Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationStorageClassifications_Get.json
     */
    /**
     * Sample code: Gets the details of a storage classification.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAStorageClassification(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.storageClassifications().getWithResponse("resourceGroupPS1", "vault1",
            "2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0", "8891569e-aaef-4a46-a4a0-78c14f2d7b09",
            com.azure.core.util.Context.NONE);
    }
}
