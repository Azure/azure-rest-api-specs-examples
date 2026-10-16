
/**
 * Samples for EdgeMachineNetworkAdapters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineNetworkAdapters_Get.json
     */
    /**
     * Sample code: EdgeMachineNetworkAdapters_Get.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineNetworkAdaptersGet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineNetworkAdapters().getWithResponse("test-rg", "EdgeMachine01", "ethernet0",
            com.azure.core.util.Context.NONE);
    }
}
