/** Samples for ReplicationRecoveryServicesProviders Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryServicesProviders_Get.json
     */
    /**
     * Sample code: Gets the details of a recovery services provider.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfARecoveryServicesProvider(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationRecoveryServicesProviders()
            .getWithResponse(
                "vault1",
                "resourceGroupPS1",
                "cloud1",
                "241641e6-ee7b-4ee4-8141-821fadda43fa",
                com.azure.core.util.Context.NONE);
    }
}
