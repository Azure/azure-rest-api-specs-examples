/** Samples for ReplicationLogicalNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationLogicalNetworks_Get.json
     */
    /**
     * Sample code: Gets a logical network with specified server id and logical network name.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsALogicalNetworkWithSpecifiedServerIdAndLogicalNetworkName(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationLogicalNetworks()
            .getWithResponse(
                "vault1",
                "resourceGroupPS1",
                "cloud1",
                "87ab394f-165f-4aa9-bd84-b018500b4509",
                com.azure.core.util.Context.NONE);
    }
}
