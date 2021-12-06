Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DomainTopics Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/DomainTopics_Delete.json
     */
    /**
     * Sample code: DomainTopics_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().delete("examplerg", "exampledomain1", "exampledomaintopic1", Context.NONE);
    }
}
```
