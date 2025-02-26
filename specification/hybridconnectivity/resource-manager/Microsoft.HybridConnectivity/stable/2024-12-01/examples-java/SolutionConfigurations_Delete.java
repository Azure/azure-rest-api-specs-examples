
/**
 * Samples for SolutionConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionConfigurations_Delete.json
     */
    /**
     * Sample code: SolutionConfigurations_Delete.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        solutionConfigurationsDelete(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionConfigurations().deleteByResourceGroupWithResponse("ymuj", "stu",
            com.azure.core.util.Context.NONE);
    }
}
