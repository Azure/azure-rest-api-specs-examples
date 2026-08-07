
/**
 * Samples for AIManagerNamespaces RotateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagerNamespaces_RotateKeys.json
     */
    /**
     * Sample code: AIManagerNamespaces_RotateKeys.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void aIManagerNamespacesRotateKeys(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagerNamespaces().rotateKeysWithResponse("rgaimanagers", "aimanager1", "namespace-1",
            com.azure.core.util.Context.NONE);
    }
}
