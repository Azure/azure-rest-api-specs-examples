
/**
 * Samples for SolutionTypes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionTypes_Get.json
     */
    /**
     * Sample code: SolutionTypes_Get.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        solutionTypesGet(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionTypes().getByResourceGroupWithResponse("rgpublicCloud", "lulzqllpu",
            com.azure.core.util.Context.NONE);
    }
}
