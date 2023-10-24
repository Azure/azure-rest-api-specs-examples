/** Samples for VaultOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/VaultOperationStatus_Get.json
     */
    /**
     * Sample code: VaultOperationStatus_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void vaultOperationStatusGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .vaultOperationStatus()
            .getWithResponse("rgrecoveryservicesdatareplication", "4", "vsdvwe", com.azure.core.util.Context.NONE);
    }
}
