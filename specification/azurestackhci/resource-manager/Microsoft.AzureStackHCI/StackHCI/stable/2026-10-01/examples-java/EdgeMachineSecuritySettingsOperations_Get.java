
/**
 * Samples for EdgeMachineSecuritySettingsOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineSecuritySettingsOperations_Get.json
     */
    /**
     * Sample code: Get edge machine security settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        getEdgeMachineSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineSecuritySettingsOperations().getWithResponse("test-rg", "testEdgeMachine", "default",
            com.azure.core.util.Context.NONE);
    }
}
