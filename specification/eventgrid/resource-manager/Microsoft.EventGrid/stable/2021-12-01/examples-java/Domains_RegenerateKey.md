Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.4/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.DomainRegenerateKeyRequest;

/** Samples for Domains RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/Domains_RegenerateKey.json
     */
    /**
     * Sample code: Domains_RegenerateKey.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsRegenerateKey(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domains()
            .regenerateKeyWithResponse(
                "examplerg", "exampledomain2", new DomainRegenerateKeyRequest().withKeyName("key1"), Context.NONE);
    }
}
```
