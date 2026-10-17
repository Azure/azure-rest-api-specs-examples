
/**
 * Samples for EdgeMachines RefreshSecuritySettings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachines_RefreshSecuritySettings.json
     */
    /**
     * Sample code: Refresh edge machine security settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        refreshEdgeMachineSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachines().refreshSecuritySettings("test-rg", "testEdgeMachine", com.azure.core.util.Context.NONE);
    }
}
