
/**
 * Samples for Vault List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Vault_ListBySubscription.json
     */
    /**
     * Sample code: Lists the vaults by subscription.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheVaultsBySubscription(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.vaults().list(com.azure.core.util.Context.NONE);
    }
}
