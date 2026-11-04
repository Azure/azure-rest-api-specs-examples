
/**
 * Samples for ReplicationVaultHealth Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationVaultHealth_Get.json
     */
    /**
     * Sample code: Gets the health summary for the vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheHealthSummaryForTheVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationVaultHealths().getWithResponse("resourceGroupPS1", "vault1",
            com.azure.core.util.Context.NONE);
    }
}
