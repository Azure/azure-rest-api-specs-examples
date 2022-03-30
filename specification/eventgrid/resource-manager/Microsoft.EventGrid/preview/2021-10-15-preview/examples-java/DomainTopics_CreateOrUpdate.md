Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.1/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DomainTopics CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainTopics_CreateOrUpdate.json
     */
    /**
     * Sample code: DomainTopics_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().createOrUpdate("examplerg", "exampledomain1", "exampledomaintopic1", Context.NONE);
    }
}
```
