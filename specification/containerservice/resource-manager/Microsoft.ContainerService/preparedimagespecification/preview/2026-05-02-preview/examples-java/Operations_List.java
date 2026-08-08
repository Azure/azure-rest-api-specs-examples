
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/Operations_List.json
     */
    /**
     * Sample code: List the operations for the provider.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void listTheOperationsForTheProvider(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
