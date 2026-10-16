
/**
 * Samples for EdgeMachineVolumes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineVolumes_List.json
     */
    /**
     * Sample code: EdgeMachineVolumes_List.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineVolumesList(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineVolumes().list("test-rg", "EdgeMachine01", com.azure.core.util.Context.NONE);
    }
}
