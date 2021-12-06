Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.eventhubs.models.ConnectionState;
import com.azure.resourcemanager.eventhubs.models.EndPointProvisioningState;
import com.azure.resourcemanager.eventhubs.models.PrivateEndpoint;
import com.azure.resourcemanager.eventhubs.models.PrivateLinkConnectionStatus;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionCreate.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpacePrivateEndPointConnectionCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .createOrUpdateWithResponse(
                "ArunMonocle",
                "sdk-Namespace-2924",
                "privateEndpointConnectionName",
                new PrivateEndpointConnectionInner()
                    .withPrivateEndpoint(
                        new PrivateEndpoint()
                            .withId(
                                "/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-8396/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-2847"))
                    .withPrivateLinkServiceConnectionState(
                        new ConnectionState()
                            .withStatus(PrivateLinkConnectionStatus.REJECTED)
                            .withDescription("testing"))
                    .withProvisioningState(EndPointProvisioningState.SUCCEEDED),
                Context.NONE);
    }
}
```
