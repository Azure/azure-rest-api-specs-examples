Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PartnerTopics Deactivate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerTopics_Deactivate.json
     */
    /**
     * Sample code: PartnerTopics_Deactivate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsDeactivate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().deactivateWithResponse("examplerg", "examplePartnerTopic1", Context.NONE);
    }
}
```
