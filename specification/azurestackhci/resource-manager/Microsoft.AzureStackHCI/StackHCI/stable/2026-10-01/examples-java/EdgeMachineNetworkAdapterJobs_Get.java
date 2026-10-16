
/**
 * Samples for EdgeMachineNetworkAdapterJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineNetworkAdapterJobs_Get.json
     */
    /**
     * Sample code: EdgeMachineNetworkAdapterJobs_Get.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineNetworkAdapterJobsGet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineNetworkAdapterJobs().getWithResponse("test-rg", "EdgeMachine01", "ethernet0",
            "ApplyConfiguration", com.azure.core.util.Context.NONE);
    }
}
