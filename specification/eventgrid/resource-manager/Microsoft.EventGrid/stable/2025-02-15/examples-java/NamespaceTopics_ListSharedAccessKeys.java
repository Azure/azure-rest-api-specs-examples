
/**
 * Samples for NamespaceTopics ListSharedAccessKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * NamespaceTopics_ListSharedAccessKeys.json
     */
    /**
     * Sample code: NamespaceTopics_ListSharedAccessKeys.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        namespaceTopicsListSharedAccessKeys(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopics().listSharedAccessKeysWithResponse("examplerg", "examplenamespace2",
            "examplenamespacetopic2", com.azure.core.util.Context.NONE);
    }
}
