
/**
 * Samples for ManagedNetworkSettingsRule List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/ManagedNetwork/listRule.json
     */
    /**
     * Sample code: List ManagedNetworkSettingsRule.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listManagedNetworkSettingsRule(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.managedNetworkSettingsRules().list("test-rg", "aml-workspace-name", com.azure.core.util.Context.NONE);
    }
}
