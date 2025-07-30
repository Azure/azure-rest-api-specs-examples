
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateEndpointConnection_List.json
     */
    /**
     * Sample code: Lists the Private Endpoint Connections.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsThePrivateEndpointConnections(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateEndpointConnections().list("rgswagger_2024-09-01", "4", com.azure.core.util.Context.NONE);
    }
}
