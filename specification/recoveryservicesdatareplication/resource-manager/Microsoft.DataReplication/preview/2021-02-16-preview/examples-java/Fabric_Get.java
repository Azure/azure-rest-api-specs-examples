/** Samples for Fabric GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Fabric_Get.json
     */
    /**
     * Sample code: Fabric_Get.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void fabricGet(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .fabrics()
            .getByResourceGroupWithResponse(
                "rgrecoveryservicesdatareplication", "wPR", com.azure.core.util.Context.NONE);
    }
}
