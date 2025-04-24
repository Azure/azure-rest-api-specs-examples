
/**
 * Samples for ReplicationVaultSetting Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationVaultSetting_Get.json
     */
    /**
     * Sample code: Gets the vault setting.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheVaultSetting(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationVaultSettings().getWithResponse("resourceGroupPS1", "vault1", "default",
            com.azure.core.util.Context.NONE);
    }
}
