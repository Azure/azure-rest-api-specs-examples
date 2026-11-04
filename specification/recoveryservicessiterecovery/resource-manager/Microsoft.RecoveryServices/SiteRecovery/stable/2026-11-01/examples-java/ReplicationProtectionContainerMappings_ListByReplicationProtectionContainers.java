
/**
 * Samples for ReplicationProtectionContainerMappings ListByReplicationProtectionContainers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionContainerMappings_ListByReplicationProtectionContainers.json
     */
    /**
     * Sample code: Gets the list of protection container mappings for a protection container.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfProtectionContainerMappingsForAProtectionContainer(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainerMappings().listByReplicationProtectionContainers("resourceGroupPS1",
            "vault1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", com.azure.core.util.Context.NONE);
    }
}
