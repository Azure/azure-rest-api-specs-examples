Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: Get Private Endpoint Connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .getWithResponse("rg1", "clustername1", "privateendpointconnection1", Context.NONE);
    }
}
```
