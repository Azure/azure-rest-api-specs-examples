
/**
 * Samples for Registries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registries/get-UserCreated.json
     */
    /**
     * Sample code: Get Registry with user created accounts.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryWithUserCreatedAccounts(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registries().getByResourceGroupWithResponse("test-rg", "string", com.azure.core.util.Context.NONE);
    }
}
