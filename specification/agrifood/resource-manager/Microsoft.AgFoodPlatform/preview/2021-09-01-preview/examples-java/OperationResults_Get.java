
/**
 * Samples for OperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/
     * OperationResults_Get.json
     */
    /**
     * Sample code: OperationResults_Get.
     * 
     * @param manager Entry point to AgriFoodManager.
     */
    public static void operationResultsGet(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.operationResults().getWithResponse("westus", "resource-provisioning-id-farmBeatsResourceName",
            com.azure.core.util.Context.NONE);
    }
}
