
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Operations_List.json
     */
    /**
     * Sample code: List all operations.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listAllOperations(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
