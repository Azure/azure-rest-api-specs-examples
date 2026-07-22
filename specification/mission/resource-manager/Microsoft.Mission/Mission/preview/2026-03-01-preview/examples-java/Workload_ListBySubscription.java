
/**
 * Samples for Workload ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Workload_ListBySubscription.json
     */
    /**
     * Sample code: Workload_ListBySubscription.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void workloadListBySubscription(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.workloads().listBySubscription("TestMyEnclave", com.azure.core.util.Context.NONE);
    }
}
