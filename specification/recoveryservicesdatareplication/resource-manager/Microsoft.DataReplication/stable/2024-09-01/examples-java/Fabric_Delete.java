
/**
 * Samples for Fabric Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Fabric_Delete.json
     */
    /**
     * Sample code: Deletes the fabric.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesTheFabric(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabrics().delete("rgrecoveryservicesdatareplication", "wPR", com.azure.core.util.Context.NONE);
    }
}
