
import com.azure.resourcemanager.recoveryservicesdatareplication.models.VaultModel;

/**
 * Samples for Vault Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Vault_Update.json
     */
    /**
     * Sample code: Updates the vault.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void updatesTheVault(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        VaultModel resource = manager.vaults()
            .getByResourceGroupWithResponse("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
