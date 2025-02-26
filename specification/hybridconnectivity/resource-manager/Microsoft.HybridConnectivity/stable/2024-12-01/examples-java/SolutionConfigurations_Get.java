
/**
 * Samples for SolutionConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionConfigurations_Get.json
     */
    /**
     * Sample code: SolutionConfigurations_Get.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        solutionConfigurationsGet(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionConfigurations().getWithResponse("ymuj", "tks", com.azure.core.util.Context.NONE);
    }
}
