Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ConsumerGroups ListByEventHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/ConsumerGroup/EHConsumerGroupListByEventHub.json
     */
    /**
     * Sample code: ConsumerGroupsListAll.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void consumerGroupsListAll(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getConsumerGroups()
            .listByEventHub("ArunMonocle", "sdk-Namespace-2661", "sdk-EventHub-6681", null, null, Context.NONE);
    }
}
```
