
/**
 * Samples for NamespaceTopics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/NamespaceTopics_Delete.
     * json
     */
    /**
     * Sample code: NamespaceTopics_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopics().delete("examplerg", "examplenamespace2", "examplenamespacetopic2",
            com.azure.core.util.Context.NONE);
    }
}
