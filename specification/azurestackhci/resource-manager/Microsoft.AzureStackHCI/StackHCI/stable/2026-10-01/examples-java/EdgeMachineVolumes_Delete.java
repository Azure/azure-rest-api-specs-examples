
/**
 * Samples for EdgeMachineVolumes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineVolumes_Delete.json
     */
    /**
     * Sample code: EdgeMachineVolumes_Delete.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineVolumesDelete(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineVolumes().delete("test-rg", "EdgeMachine01", "vol-001", com.azure.core.util.Context.NONE);
    }
}
