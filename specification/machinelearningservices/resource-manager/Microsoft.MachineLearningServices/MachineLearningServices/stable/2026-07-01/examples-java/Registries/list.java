
/**
 * Samples for Registries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Registries/list.json
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
