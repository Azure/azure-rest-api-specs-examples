import com.azure.core.util.Context;

/** Samples for FarmBeatsModels List. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_ListBySubscription.json
     */
    /**
     * Sample code: FarmBeatsModels_ListBySubscription.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void farmBeatsModelsListBySubscription(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.farmBeatsModels().list(null, null, Context.NONE);
    }
}
