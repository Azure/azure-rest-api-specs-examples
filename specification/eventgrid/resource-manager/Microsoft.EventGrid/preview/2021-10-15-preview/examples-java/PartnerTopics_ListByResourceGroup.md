Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.1/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PartnerTopics ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerTopics_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerTopics_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
```
