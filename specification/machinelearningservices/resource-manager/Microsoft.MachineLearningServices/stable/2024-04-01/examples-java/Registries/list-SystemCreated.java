
/**
 * Samples for Registries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registries/list-SystemCreated.json
     */
    /**
     * Sample code: List registries with system created accounts.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listRegistriesWithSystemCreatedAccounts(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registries().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
