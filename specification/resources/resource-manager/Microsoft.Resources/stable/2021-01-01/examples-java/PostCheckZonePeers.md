Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.resources.models.CheckZonePeersRequest;
import java.util.Arrays;

/** Samples for Subscriptions CheckZonePeers. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PostCheckZonePeers.json
     */
    /**
     * Sample code: GetLogicalZoneMapping.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLogicalZoneMapping(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .subscriptionClient()
            .getSubscriptions()
            .checkZonePeersWithResponse(
                "00000000-0000-0000-0000-00000000000000",
                new CheckZonePeersRequest()
                    .withLocation("eastus")
                    .withSubscriptionIds(Arrays.asList("subscriptions/11111111-1111-1111-1111-111111111111")),
                Context.NONE);
    }
}
```
