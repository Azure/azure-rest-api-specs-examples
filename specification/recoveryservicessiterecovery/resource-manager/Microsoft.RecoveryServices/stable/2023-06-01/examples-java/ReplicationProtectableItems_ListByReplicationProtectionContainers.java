/** Samples for ReplicationProtectableItems ListByReplicationProtectionContainers. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectableItems_ListByReplicationProtectionContainers.json
     */
    /**
     * Sample code: Gets the list of protectable items.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfProtectableItems(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationProtectableItems()
            .listByReplicationProtectionContainers(
                "vault1",
                "resourceGroupPS1",
                "cloud1",
                "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
