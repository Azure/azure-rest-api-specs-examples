
/**
 * Samples for EdgeDevices Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * GetEdgeDevices.json
     */
    /**
     * Sample code: Get Edge Device.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getEdgeDevice(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeDevices().getWithResponse(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1",
            "default", com.azure.core.util.Context.NONE);
    }
}
