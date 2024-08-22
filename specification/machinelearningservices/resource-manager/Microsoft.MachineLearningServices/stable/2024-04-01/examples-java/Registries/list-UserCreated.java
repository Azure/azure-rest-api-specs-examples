
/**
 * Samples for Registries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registries/list-UserCreated.json
     */
    /**
     * Sample code: List registries with user created accounts.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listRegistriesWithUserCreatedAccounts(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registries().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
