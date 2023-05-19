/** Samples for NamespaceTopics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/NamespaceTopics_Get.json
     */
    /**
     * Sample code: NamespaceTopics_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .namespaceTopics()
            .getWithResponse(
                "examplerg", "examplenamespace2", "examplenamespacetopic2", com.azure.core.util.Context.NONE);
    }
}
