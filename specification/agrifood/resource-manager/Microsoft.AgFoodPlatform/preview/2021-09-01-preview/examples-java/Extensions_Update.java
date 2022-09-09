import com.azure.core.util.Context;

/** Samples for Extensions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Extensions_Update.json
     */
    /**
     * Sample code: Extensions_Update.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void extensionsUpdate(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .extensions()
            .updateWithResponse("examples-rg", "examples-farmbeatsResourceName", "provider.extension", Context.NONE);
    }
}
