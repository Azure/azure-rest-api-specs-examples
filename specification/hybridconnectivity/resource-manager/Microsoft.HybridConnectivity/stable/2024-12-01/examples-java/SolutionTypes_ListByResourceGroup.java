
/**
 * Samples for SolutionTypes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionTypes_ListByResourceGroup.json
     */
    /**
     * Sample code: SolutionTypes_ListByResourceGroup.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void solutionTypesListByResourceGroup(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionTypes().listByResourceGroup("rgpublicCloud", com.azure.core.util.Context.NONE);
    }
}
