
/**
 * Samples for EdgeMachineNetworkAdapterJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineNetworkAdapterJobs_Delete.json
     */
    /**
     * Sample code: EdgeMachineNetworkAdapterJobs_Delete.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineNetworkAdapterJobsDelete(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineNetworkAdapterJobs().delete("test-rg", "EdgeMachine01", "ethernet0", "ApplyConfiguration",
            com.azure.core.util.Context.NONE);
    }
}
