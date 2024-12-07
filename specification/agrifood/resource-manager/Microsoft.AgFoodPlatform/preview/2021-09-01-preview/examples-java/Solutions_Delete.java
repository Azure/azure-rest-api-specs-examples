
/**
 * Samples for Solutions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/
     * Solutions_Delete.json
     */
    /**
     * Sample code: Solutions_Delete.
     * 
     * @param manager Entry point to AgriFoodManager.
     */
    public static void solutionsDelete(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.solutions().deleteWithResponse("examples-rg", "examples-farmbeatsResourceName", "provider.solution",
            com.azure.core.util.Context.NONE);
    }
}
