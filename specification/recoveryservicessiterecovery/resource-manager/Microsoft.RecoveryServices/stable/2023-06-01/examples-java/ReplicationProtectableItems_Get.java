/** Samples for ReplicationProtectableItems Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectableItems_Get.json
     */
    /**
     * Sample code: Gets the details of a protectable item.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAProtectableItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationProtectableItems()
            .getWithResponse(
                "vault1",
                "resourceGroupPS1",
                "cloud1",
                "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
                "c0c14913-3d7a-48ea-9531-cc99e0e686e6",
                com.azure.core.util.Context.NONE);
    }
}
