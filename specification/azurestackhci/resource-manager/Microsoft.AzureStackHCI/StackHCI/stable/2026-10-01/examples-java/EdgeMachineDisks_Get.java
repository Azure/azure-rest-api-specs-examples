
/**
 * Samples for EdgeMachineDisks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDisks_Get.json
     */
    /**
     * Sample code: EdgeMachineDisks_Get.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineDisksGet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDisks().getWithResponse("test-rg", "EdgeMachine01", "disk-001",
            com.azure.core.util.Context.NONE);
    }
}
