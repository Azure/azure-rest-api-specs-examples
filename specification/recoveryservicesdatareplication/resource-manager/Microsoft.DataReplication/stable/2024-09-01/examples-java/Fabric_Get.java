
/**
 * Samples for Fabric GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Fabric_Get.json
     */
    /**
     * Sample code: Gets the fabric.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheFabric(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabrics().getByResourceGroupWithResponse("rgrecoveryservicesdatareplication", "wPR",
            com.azure.core.util.Context.NONE);
    }
}
