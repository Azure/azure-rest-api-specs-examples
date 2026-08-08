
/**
 * Samples for AIManagerNamespaces ListAccessKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagerNamespaces_ListAccessKeys.json
     */
    /**
     * Sample code: AIManagerNamespaces_ListAccessKeys_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void aIManagerNamespacesListAccessKeysMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagerNamespaces().listAccessKeysWithResponse("rgaimanagers", "aimanager1", "namespace-1",
            com.azure.core.util.Context.NONE);
    }
}
