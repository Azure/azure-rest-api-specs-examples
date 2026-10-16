
import com.azure.resourcemanager.azurestackhci.models.EdgeMachineSecurityProperties;

/**
 * Samples for EdgeMachineSecuritySettingsOperations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineSecuritySettingsOperations_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update edge machine security settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createOrUpdateEdgeMachineSecuritySettings(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineSecuritySettingsOperations().define("default")
            .withExistingEdgeMachine("test-rg", "testEdgeMachine").withProperties(new EdgeMachineSecurityProperties())
            .create();
    }
}
