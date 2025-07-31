
/**
 * Samples for Vault Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Vault_Delete.json
     */
    /**
     * Sample code: Deletes the vault.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesTheVault(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.vaults().delete("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
