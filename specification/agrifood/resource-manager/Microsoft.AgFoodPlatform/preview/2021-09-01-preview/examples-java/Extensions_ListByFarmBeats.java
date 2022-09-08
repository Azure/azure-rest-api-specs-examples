import com.azure.core.util.Context;

/** Samples for Extensions ListByFarmBeats. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Extensions_ListByFarmBeats.json
     */
    /**
     * Sample code: Extensions_ListByFarmBeats.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void extensionsListByFarmBeats(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .extensions()
            .listByFarmBeats("examples-rg", "examples-farmbeatsResourceName", null, null, null, null, Context.NONE);
    }
}
