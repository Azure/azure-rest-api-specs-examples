
/**
 * Samples for ReplicationVaultHealth Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationVaultHealth_Refresh.json
     */
    /**
     * Sample code: Refreshes health summary of the vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void refreshesHealthSummaryOfTheVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationVaultHealths().refresh("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
