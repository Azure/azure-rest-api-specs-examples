
/**
 * Samples for EdgeMachineVolumes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineVolumes_Get.json
     */
    /**
     * Sample code: EdgeMachineVolumes_Get.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineVolumesGet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineVolumes().getWithResponse("test-rg", "EdgeMachine01", "vol-001",
            com.azure.core.util.Context.NONE);
    }
}
