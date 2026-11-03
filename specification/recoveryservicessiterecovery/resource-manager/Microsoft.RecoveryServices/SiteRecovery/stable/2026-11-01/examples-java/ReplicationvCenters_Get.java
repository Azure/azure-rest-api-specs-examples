
/**
 * Samples for ReplicationvCenters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationvCenters_Get.json
     */
    /**
     * Sample code: Gets the details of a vCenter.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheDetailsOfAVCenter(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationvCenters().getWithResponse("MadhaviVRG", "MadhaviVault", "MadhaviFabric", "esx-78",
            com.azure.core.util.Context.NONE);
    }
}
