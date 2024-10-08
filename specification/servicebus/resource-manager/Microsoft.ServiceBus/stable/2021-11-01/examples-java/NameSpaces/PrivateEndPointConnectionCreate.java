
import com.azure.resourcemanager.servicebus.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.servicebus.models.ConnectionState;
import com.azure.resourcemanager.servicebus.models.EndPointProvisioningState;
import com.azure.resourcemanager.servicebus.models.PrivateEndpoint;
import com.azure.resourcemanager.servicebus.models.PrivateLinkConnectionStatus;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * PrivateEndPointConnectionCreate.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpacePrivateEndPointConnectionCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getPrivateEndpointConnections()
            .createOrUpdateWithResponse("ArunMonocle", "sdk-Namespace-2924", "privateEndpointConnectionName",
                new PrivateEndpointConnectionInner().withPrivateEndpoint(new PrivateEndpoint().withId(
                    "/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-8396/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-2847"))
                    .withPrivateLinkServiceConnectionState(new ConnectionState()
                        .withStatus(PrivateLinkConnectionStatus.REJECTED).withDescription("testing"))
                    .withProvisioningState(EndPointProvisioningState.SUCCEEDED),
                com.azure.core.util.Context.NONE);
    }
}
