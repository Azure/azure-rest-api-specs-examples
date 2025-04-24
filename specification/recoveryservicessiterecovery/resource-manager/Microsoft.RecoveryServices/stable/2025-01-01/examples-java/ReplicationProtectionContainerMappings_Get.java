
/**
 * Samples for ReplicationProtectionContainerMappings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionContainerMappings_Get.json
     */
    /**
     * Sample code: Gets a protection container mapping.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsAProtectionContainerMapping(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainerMappings().getWithResponse("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "cloud1protectionprofile1", com.azure.core.util.Context.NONE);
    }
}
