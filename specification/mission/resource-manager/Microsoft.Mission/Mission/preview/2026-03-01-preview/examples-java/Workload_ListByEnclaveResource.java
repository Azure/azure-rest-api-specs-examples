
/**
 * Samples for Workload ListByEnclaveResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Workload_ListByEnclaveResource.json
     */
    /**
     * Sample code: Workload_ListByEnclaveResource.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void workloadListByEnclaveResource(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.workloads().listByEnclaveResource("rgopenapi", "TestMyEnclave", com.azure.core.util.Context.NONE);
    }
}
