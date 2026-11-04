
/**
 * Samples for ReplicationProtectionContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionContainers_Get.json
     */
    /**
     * Sample code: Gets the protection container details.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheProtectionContainerDetails(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainers().getWithResponse("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", com.azure.core.util.Context.NONE);
    }
}
