Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PartnerNamespaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerNamespaces_Get.json
     */
    /**
     * Sample code: PartnerNamespaces_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerNamespacesGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerNamespaces()
            .getByResourceGroupWithResponse("examplerg", "examplePartnerNamespaceName1", Context.NONE);
    }
}
```
