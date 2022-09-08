import com.azure.core.util.Context;

/** Samples for FarmBeatsModels ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_ListByResourceGroup.json
     */
    /**
     * Sample code: FarmBeatsModels_ListByResourceGroup.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void farmBeatsModelsListByResourceGroup(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.farmBeatsModels().listByResourceGroup("examples-rg", null, null, Context.NONE);
    }
}
