
/**
 * Samples for Workload Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Workload_Get.json
     */
    /**
     * Sample code: Workload_Get.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void workloadGet(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.workloads().getWithResponse("rgopenapi", "TestMyEnclave", "TestMyWorkload",
            com.azure.core.util.Context.NONE);
    }
}
