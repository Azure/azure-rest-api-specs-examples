
/**
 * Samples for FabricAgent Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/FabricAgent_Delete.json
     */
    /**
     * Sample code: Deletes the Fabric Agent.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesTheFabricAgent(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabricAgents().delete("rgswagger_2024-09-01", "wPR", "M", com.azure.core.util.Context.NONE);
    }
}
