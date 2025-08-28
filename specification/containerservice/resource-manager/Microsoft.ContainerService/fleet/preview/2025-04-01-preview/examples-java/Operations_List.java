
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/Operations_List.json
     */
    /**
     * Sample code: List the operations for the provider.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listTheOperationsForTheProvider(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
