
/**
 * Samples for EdgeMachineNetworkAdapters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineNetworkAdapters_List.json
     */
    /**
     * Sample code: EdgeMachineNetworkAdapters_List.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineNetworkAdaptersList(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineNetworkAdapters().list("test-rg", "EdgeMachine01", com.azure.core.util.Context.NONE);
    }
}
