Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.PartnerNamespaceRegenerateKeyRequest;

/** Samples for PartnerNamespaces RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerNamespaces_RegenerateKey.json
     */
    /**
     * Sample code: PartnerNamespaces_RegenerateKey.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerNamespacesRegenerateKey(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerNamespaces()
            .regenerateKeyWithResponse(
                "examplerg",
                "examplePartnerNamespaceName1",
                new PartnerNamespaceRegenerateKeyRequest().withKeyName("key1"),
                Context.NONE);
    }
}
```
