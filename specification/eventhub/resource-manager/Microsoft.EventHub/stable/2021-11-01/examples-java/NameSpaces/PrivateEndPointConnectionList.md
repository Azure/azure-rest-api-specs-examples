Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionList.json
     */
    /**
     * Sample code: NameSpaceCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .list("SDK-EventHub-4794", "sdk-Namespace-5828", Context.NONE);
    }
}
```
