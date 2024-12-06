
/**
 * Samples for SolutionsDiscoverability List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/
     * SolutionsDiscoverability_List.json
     */
    /**
     * Sample code: SolutionsDiscoverability_List.
     * 
     * @param manager Entry point to AgriFoodManager.
     */
    public static void solutionsDiscoverabilityList(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.solutionsDiscoverabilities().list(null, null, null, com.azure.core.util.Context.NONE);
    }
}
