
/**
 * Samples for Vault GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Vault_Get.json
     */
    /**
     * Sample code: Gets the vault.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheVault(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.vaults().getByResourceGroupWithResponse("rgrecoveryservicesdatareplication", "4",
            com.azure.core.util.Context.NONE);
    }
}
