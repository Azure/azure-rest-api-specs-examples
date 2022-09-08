import com.azure.core.util.Context;

/** Samples for FarmBeatsExtensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsExtensions_List.json
     */
    /**
     * Sample code: FarmBeatsExtensions_List.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void farmBeatsExtensionsList(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.farmBeatsExtensions().list(null, null, null, null, null, Context.NONE);
    }
}
