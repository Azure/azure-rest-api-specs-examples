
/**
 * Samples for FabricAgent List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/FabricAgent_List.json
     */
    /**
     * Sample code: Lists the fabric agents.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheFabricAgents(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabricAgents().list("rgswagger_2024-09-01", "wPR", com.azure.core.util.Context.NONE);
    }
}
