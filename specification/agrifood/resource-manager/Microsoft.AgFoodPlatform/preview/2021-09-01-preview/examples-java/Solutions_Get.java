
/**
 * Samples for Solutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/
     * Solutions_Get.json
     */
    /**
     * Sample code: Solutions_Get.
     * 
     * @param manager Entry point to AgriFoodManager.
     */
    public static void solutionsGet(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.solutions().getWithResponse("examples-rg", "examples-farmbeatsResourceName", "provider.solution",
            com.azure.core.util.Context.NONE);
    }
}
