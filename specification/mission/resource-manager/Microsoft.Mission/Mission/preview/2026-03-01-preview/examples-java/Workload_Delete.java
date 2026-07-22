
/**
 * Samples for Workload Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Workload_Delete.json
     */
    /**
     * Sample code: Workload_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void workloadDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.workloads().delete("rgopenapi", "TestMyEnclave", "TestMyWorkload", com.azure.core.util.Context.NONE);
    }
}
