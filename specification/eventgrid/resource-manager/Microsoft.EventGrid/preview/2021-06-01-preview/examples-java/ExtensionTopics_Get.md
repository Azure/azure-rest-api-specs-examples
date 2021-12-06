Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ExtensionTopics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/ExtensionTopics_Get.json
     */
    /**
     * Sample code: ExtensionTopics_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void extensionTopicsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .extensionTopics()
            .getWithResponse(
                "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/microsoft.storage/storageaccounts/exampleResourceName/providers/Microsoft.eventgrid/extensionTopics/default",
                Context.NONE);
    }
}
```
