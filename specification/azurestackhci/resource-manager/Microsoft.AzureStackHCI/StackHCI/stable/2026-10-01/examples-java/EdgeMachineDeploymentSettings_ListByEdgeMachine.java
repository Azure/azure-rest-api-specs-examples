
/**
 * Samples for EdgeMachineDeploymentSettings ListByEdgeMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDeploymentSettings_ListByEdgeMachine.json
     */
    /**
     * Sample code: EdgeMachineDeploymentSettings_ListByEdgeMachine.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineDeploymentSettingsListByEdgeMachine(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDeploymentSettings().listByEdgeMachine("test-rg", "EdgeMachine01",
            com.azure.core.util.Context.NONE);
    }
}
