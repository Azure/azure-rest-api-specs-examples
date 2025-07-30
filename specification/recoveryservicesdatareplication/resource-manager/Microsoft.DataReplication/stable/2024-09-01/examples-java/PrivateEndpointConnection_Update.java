
import com.azure.resourcemanager.recoveryservicesdatareplication.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PrivateEndpoint;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PrivateEndpointConnectionResponseProperties;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PrivateEndpointConnectionStatus;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateEndpointConnection_Update.json
     */
    /**
     * Sample code: Updates the Private Endpoint Connection.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void updatesThePrivateEndpointConnection(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateEndpointConnections().updateWithResponse("rgswagger_2024-09-01", "4", "jitf",
            new PrivateEndpointConnectionInner().withProperties(new PrivateEndpointConnectionResponseProperties()
                .withPrivateEndpoint(new PrivateEndpoint().withId("cwcdqoynostmqwdwy"))
                .withPrivateLinkServiceConnectionState(
                    new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointConnectionStatus.APPROVED)
                        .withDescription("y").withActionsRequired("afwbq"))),
            com.azure.core.util.Context.NONE);
    }
}
