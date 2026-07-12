
/**
 * Samples for EdgeDeviceJobs ListByEdgeDevice.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/EdgeDeviceJobs_ListByEdgeDevice.json
     */
    /**
     * Sample code: EdgeDeviceJobs_ListByEdgeDevice.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeDeviceJobsListByEdgeDevice(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeDeviceJobs().listByEdgeDevice(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1",
            "YE-855IEIN585-", com.azure.core.util.Context.NONE);
    }
}
