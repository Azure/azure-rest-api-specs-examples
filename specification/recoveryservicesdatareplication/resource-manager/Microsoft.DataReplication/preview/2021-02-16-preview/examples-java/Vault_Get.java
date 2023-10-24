/** Samples for Vault GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Vault_Get.json
     */
    /**
     * Sample code: Vault_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void vaultGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .vaults()
            .getByResourceGroupWithResponse("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
