
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/operations/preview/2023-10-01-preview/examples/
     * Operations_List.json
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
