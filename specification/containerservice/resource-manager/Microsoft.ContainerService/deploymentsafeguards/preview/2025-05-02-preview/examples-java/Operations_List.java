
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-02-preview/Operations_List.json
     */
    /**
     * Sample code: List the operations for the provider.
     * 
     * @param manager Entry point to ContainerServiceSafeguardsManager.
     */
    public static void listTheOperationsForTheProvider(
        com.azure.resourcemanager.containerservicesafeguards.ContainerServiceSafeguardsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
