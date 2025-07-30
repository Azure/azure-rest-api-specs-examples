
/**
 * Samples for FabricAgent Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/FabricAgent_Get.json
     */
    /**
     * Sample code: Gets the fabric agent.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheFabricAgent(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabricAgents().getWithResponse("rgrecoveryservicesdatareplication", "wPR", "M",
            com.azure.core.util.Context.NONE);
    }
}
