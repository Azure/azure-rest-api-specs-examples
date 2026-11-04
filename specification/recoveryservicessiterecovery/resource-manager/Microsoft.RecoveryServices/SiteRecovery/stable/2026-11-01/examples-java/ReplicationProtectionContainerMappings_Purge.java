
/**
 * Samples for ReplicationProtectionContainerMappings Purge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionContainerMappings_Purge.json
     */
    /**
     * Sample code: Purge protection container mapping.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void purgeProtectionContainerMapping(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainerMappings().purge("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "cloud1protectionprofile1", com.azure.core.util.Context.NONE);
    }
}
