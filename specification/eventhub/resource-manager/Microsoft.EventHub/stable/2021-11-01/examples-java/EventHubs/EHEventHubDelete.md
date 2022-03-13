Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EventHubs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubDelete.json
     */
    /**
     * Sample code: EventHubDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getEventHubs()
            .deleteWithResponse("ArunMonocle", "sdk-Namespace-5357", "sdk-EventHub-6547", Context.NONE);
    }
}
```
