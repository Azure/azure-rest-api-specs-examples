import com.azure.core.util.Context;

/** Samples for Extensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Extensions_Delete.json
     */
    /**
     * Sample code: Extensions_Delete.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void extensionsDelete(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .extensions()
            .deleteWithResponse("examples-rg", "examples-farmbeatsResourceName", "provider.extension", Context.NONE);
    }
}
