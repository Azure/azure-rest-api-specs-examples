
/**
 * Samples for EdgeMachineNetworkAdapterJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineNetworkAdapterJobs_List.json
     */
    /**
     * Sample code: EdgeMachineNetworkAdapterJobs_List.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineNetworkAdapterJobsList(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineNetworkAdapterJobs().list("test-rg", "EdgeMachine01", "ethernet0",
            com.azure.core.util.Context.NONE);
    }
}
