
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateEndpointConnection_Delete.json
     */
    /**
     * Sample code: Deletes the Private Endpoint Connection.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesThePrivateEndpointConnection(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateEndpointConnections().delete("rgswagger_2024-09-01", "4", "sdwqtfhigjirrzhpbmqtzgs",
            com.azure.core.util.Context.NONE);
    }
}
