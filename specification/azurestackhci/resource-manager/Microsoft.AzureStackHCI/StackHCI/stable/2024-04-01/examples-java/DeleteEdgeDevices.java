
/**
 * Samples for EdgeDevices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * DeleteEdgeDevices.json
     */
    /**
     * Sample code: Delete Edge Devices.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteEdgeDevices(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeDevices().delete(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1",
            "default", com.azure.core.util.Context.NONE);
    }
}
