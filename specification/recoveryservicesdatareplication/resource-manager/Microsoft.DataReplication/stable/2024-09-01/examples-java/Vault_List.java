
/**
 * Samples for Vault ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Vault_List.json
     */
    /**
     * Sample code: Lists the vaults.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheVaults(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.vaults().listByResourceGroup("rgrecoveryservicesdatareplication", "mwculdaqndp",
            com.azure.core.util.Context.NONE);
    }
}
