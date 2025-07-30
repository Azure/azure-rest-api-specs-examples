
/**
 * Samples for Vault Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Vault_Create.json
     */
    /**
     * Sample code: Puts the vault.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void putsTheVault(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.vaults().define("4").withRegion((String) null)
            .withExistingResourceGroup("rgrecoveryservicesdatareplication").create();
    }
}
