
/**
 * Samples for EdgeActions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActions_ListBySubscription.json
     */
    /**
     * Sample code: ListEdgeActions_bySubscription.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void listEdgeActionsBySubscription(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActions().list(com.azure.core.util.Context.NONE);
    }
}
