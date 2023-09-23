/** Samples for ReplicationEvents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationEvents_Get.json
     */
    /**
     * Sample code: Get the details of an Azure Site recovery event.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getTheDetailsOfAnAzureSiteRecoveryEvent(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationEvents()
            .getWithResponse(
                "vault1", "resourceGroupPS1", "654b71d0-b2ce-4e6e-a861-98528d4bd375", com.azure.core.util.Context.NONE);
    }
}
