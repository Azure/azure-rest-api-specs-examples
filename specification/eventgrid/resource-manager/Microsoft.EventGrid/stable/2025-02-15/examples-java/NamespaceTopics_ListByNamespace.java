
/**
 * Samples for NamespaceTopics ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * NamespaceTopics_ListByNamespace.json
     */
    /**
     * Sample code: NamespaceTopics_ListByNamespace.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicsListByNamespace(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopics().listByNamespace("examplerg", "examplenamespace2", null, null,
            com.azure.core.util.Context.NONE);
    }
}
