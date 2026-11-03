
/**
 * Samples for ReplicationProtectionContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionContainers_List.json
     */
    /**
     * Sample code: Gets the list of all protection containers in a vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfAllProtectionContainersInAVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainers().list("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
