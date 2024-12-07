
/**
 * Samples for SolutionsDiscoverability Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/
     * SolutionsDiscoverability_Get.json
     */
    /**
     * Sample code: SolutionsDiscoverability_Get.
     * 
     * @param manager Entry point to AgriFoodManager.
     */
    public static void solutionsDiscoverabilityGet(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.solutionsDiscoverabilities().getWithResponse("bayerAgPowered.gdu", com.azure.core.util.Context.NONE);
    }
}
