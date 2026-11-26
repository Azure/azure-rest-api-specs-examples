
/**
 * Samples for Registries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registries/get.json
     */
    /**
     * Sample code: Get Registry with system created accounts.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getRegistryWithSystemCreatedAccounts(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registries().getByResourceGroupWithResponse("test-rg", "string", com.azure.core.util.Context.NONE);
    }
}
