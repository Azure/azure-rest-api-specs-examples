
/**
 * Samples for EdgeMachineDeploymentSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDeploymentSettings_Delete.json
     */
    /**
     * Sample code: EdgeMachineDeploymentSettings_Delete.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineDeploymentSettingsDelete(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDeploymentSettings().delete("test-rg", "EdgeMachine01", "default",
            com.azure.core.util.Context.NONE);
    }
}
