Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.3/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListPrivateEndpointConnections.json
     */
    /**
     * Sample code: List Private Endpoint Connections.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void listPrivateEndpointConnections(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.privateEndpointConnections().list("res6977", "sto2527", Context.NONE);
    }
}
```
