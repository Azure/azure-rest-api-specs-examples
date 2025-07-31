
/**
 * Samples for FabricAgent Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/FabricAgent_Create.json
     */
    /**
     * Sample code: Puts the fabric agent.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void putsTheFabricAgent(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabricAgents().define("M").withExistingReplicationFabric("rgswagger_2024-09-01", "wPR").create();
    }
}
