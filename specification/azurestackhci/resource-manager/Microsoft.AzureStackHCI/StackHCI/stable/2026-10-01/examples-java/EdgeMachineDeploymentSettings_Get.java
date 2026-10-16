
/**
 * Samples for EdgeMachineDeploymentSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDeploymentSettings_Get.json
     */
    /**
     * Sample code: EdgeMachineDeploymentSettings_Get.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineDeploymentSettingsGet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDeploymentSettings().getWithResponse("test-rg", "EdgeMachine01", "default",
            com.azure.core.util.Context.NONE);
    }
}
