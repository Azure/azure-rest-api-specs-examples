```java
import com.azure.core.util.Context;

/** Samples for Domains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/Domains_Delete.json
     */
    /**
     * Sample code: Domains_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domains().delete("examplerg", "exampledomain1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.
