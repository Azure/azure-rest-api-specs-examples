/** Samples for ReplicationvCenters Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationvCenters_Get.json
     */
    /**
     * Sample code: Gets the details of a vCenter.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAVCenter(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationvCenters()
            .getWithResponse("MadhaviVault", "MadhaviVRG", "MadhaviFabric", "esx-78", com.azure.core.util.Context.NONE);
    }
}
