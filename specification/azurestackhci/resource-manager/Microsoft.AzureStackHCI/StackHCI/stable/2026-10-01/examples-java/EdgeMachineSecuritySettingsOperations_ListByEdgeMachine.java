
/**
 * Samples for EdgeMachineSecuritySettingsOperations ListByEdgeMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineSecuritySettingsOperations_ListByEdgeMachine.json
     */
    /**
     * Sample code: List edge machine security settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listEdgeMachineSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineSecuritySettingsOperations().listByEdgeMachine("test-rg", "testEdgeMachine",
            com.azure.core.util.Context.NONE);
    }
}
