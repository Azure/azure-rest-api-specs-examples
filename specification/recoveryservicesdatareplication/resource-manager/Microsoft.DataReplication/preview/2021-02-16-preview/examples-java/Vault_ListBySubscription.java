/** Samples for Vault List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Vault_ListBySubscription.json
     */
    /**
     * Sample code: Vault_ListBySubscription.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void vaultListBySubscription(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.vaults().list("dqsjhseyugyexxrlrln", com.azure.core.util.Context.NONE);
    }
}
