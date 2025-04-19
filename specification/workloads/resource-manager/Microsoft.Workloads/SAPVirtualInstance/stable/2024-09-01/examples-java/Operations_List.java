
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Operations_List.json
     */
    /**
     * Sample code: List the operations for the provider.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void listTheOperationsForTheProvider(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
