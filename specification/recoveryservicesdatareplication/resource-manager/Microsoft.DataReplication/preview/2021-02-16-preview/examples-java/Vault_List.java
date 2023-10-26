/** Samples for Vault ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Vault_List.json
     */
    /**
     * Sample code: Vault_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void vaultList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .vaults()
            .listByResourceGroup("rgrecoveryservicesdatareplication", "mwculdaqndp", com.azure.core.util.Context.NONE);
    }
}
