
/**
 * Samples for AIManagerNamespaces ListByAIManager.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagerNamespaces_ListByAIManager.json
     */
    /**
     * Sample code: Lists AI Manager namespace resources by AI Manager.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void listsAIManagerNamespaceResourcesByAIManager(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagerNamespaces().listByAIManager("rg1", "aimanager1", com.azure.core.util.Context.NONE);
    }
}
