```java
import com.azure.core.util.Context;

/** Samples for DomainTopics ListByDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainTopics_ListByDomain.json
     */
    /**
     * Sample code: DomainTopics_ListByDomain.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsListByDomain(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().listByDomain("examplerg", "exampledomain2", null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.
