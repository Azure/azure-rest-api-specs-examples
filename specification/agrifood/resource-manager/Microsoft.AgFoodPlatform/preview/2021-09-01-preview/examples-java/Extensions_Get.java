import com.azure.core.util.Context;

/** Samples for Extensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Extensions_Get.json
     */
    /**
     * Sample code: Extensions_Get.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void extensionsGet(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .extensions()
            .getWithResponse("examples-rg", "examples-farmbeatsResourceName", "provider.extension", Context.NONE);
    }
}
