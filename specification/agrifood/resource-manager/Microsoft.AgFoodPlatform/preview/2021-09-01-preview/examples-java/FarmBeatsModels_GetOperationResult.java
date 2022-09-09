import com.azure.core.util.Context;

/** Samples for FarmBeatsModels GetOperationResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_GetOperationResult.json
     */
    /**
     * Sample code: FarmBeatsModels_GetOperationResult.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void farmBeatsModelsGetOperationResult(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .farmBeatsModels()
            .getOperationResultWithResponse(
                "examples-rg",
                "examples-farmBeatsResourceName",
                "resource-provisioning-id-farmBeatsResourceName",
                Context.NONE);
    }
}
