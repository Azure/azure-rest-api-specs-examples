
/**
 * Samples for EdgeMachineSecuritySettingsOperations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineSecuritySettingsOperations_Delete.json
     */
    /**
     * Sample code: Delete edge machine security settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        deleteEdgeMachineSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineSecuritySettingsOperations().delete("test-rg", "testEdgeMachine", "default",
            com.azure.core.util.Context.NONE);
    }
}
