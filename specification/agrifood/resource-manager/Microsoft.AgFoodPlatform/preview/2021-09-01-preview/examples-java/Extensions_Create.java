import com.azure.core.util.Context;

/** Samples for Extensions Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Extensions_Create.json
     */
    /**
     * Sample code: Extensions_Create.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void extensionsCreate(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .extensions()
            .createWithResponse("examples-rg", "examples-farmbeatsResourceName", "provider.extension", Context.NONE);
    }
}
