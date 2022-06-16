import com.azure.core.util.Context;

/** Samples for ExtensionTopics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/ExtensionTopics_Get.json
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
