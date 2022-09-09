import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateLinkResources_Get.json
     */
    /**
     * Sample code: PrivateLinkResources_Get.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void privateLinkResourcesGet(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .privateLinkResources()
            .getWithResponse("examples-rg", "examples-farmbeatsResourceName", "farmbeats", Context.NONE);
    }
}
