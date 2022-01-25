Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/PrivateLinkResources_ListByResource.json
     */
    /**
     * Sample code: PrivateLinkResources_ListByResource.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void privateLinkResourcesListByResource(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.privateLinkResources().listByResource("examplerg", "topics", "exampletopic1", null, null, Context.NONE);
    }
}
```
