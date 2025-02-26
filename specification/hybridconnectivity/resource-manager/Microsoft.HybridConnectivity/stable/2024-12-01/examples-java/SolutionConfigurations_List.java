
/**
 * Samples for SolutionConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionConfigurations_List.json
     */
    /**
     * Sample code: SolutionConfigurations_List.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        solutionConfigurationsList(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionConfigurations().list("ymuj", com.azure.core.util.Context.NONE);
    }
}
