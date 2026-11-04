
/**
 * Samples for ReplicationProtectionContainerMappings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionContainerMappings_List.json
     */
    /**
     * Sample code: Gets the list of all protection container mappings in a vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfAllProtectionContainerMappingsInAVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainerMappings().list("resourceGroupPS1", "vault1",
            com.azure.core.util.Context.NONE);
    }
}
