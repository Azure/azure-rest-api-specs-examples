
/**
 * Samples for ReplicationVaultSetting List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationVaultSetting_List.json
     */
    /**
     * Sample code: Gets the list of vault setting.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheListOfVaultSetting(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationVaultSettings().list("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
