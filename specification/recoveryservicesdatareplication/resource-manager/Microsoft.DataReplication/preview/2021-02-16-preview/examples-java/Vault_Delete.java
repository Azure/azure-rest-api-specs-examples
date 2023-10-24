/** Samples for Vault Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Vault_Delete.json
     */
    /**
     * Sample code: Vault_Delete.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void vaultDelete(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.vaults().delete("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
