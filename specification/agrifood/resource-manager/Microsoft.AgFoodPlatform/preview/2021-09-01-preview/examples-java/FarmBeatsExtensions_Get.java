import com.azure.core.util.Context;

/** Samples for FarmBeatsExtensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsExtensions_Get.json
     */
    /**
     * Sample code: FarmBeatsExtensions_Get.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void farmBeatsExtensionsGet(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.farmBeatsExtensions().getWithResponse("DTN.ContentServices", Context.NONE);
    }
}
