Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.containerservice.models.ConnectionStatus;
import com.azure.resourcemanager.containerservice.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/PrivateEndpointConnectionsUpdate.json
     */
    /**
     * Sample code: Update Private Endpoint Connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatePrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .updateWithResponse(
                "rg1",
                "clustername1",
                "privateendpointconnection1",
                new PrivateEndpointConnectionInner()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState().withStatus(ConnectionStatus.APPROVED)),
                Context.NONE);
    }
}
```
