
/**
 * Samples for Extensions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/
     * Extensions_CreateOrUpdate.json
     */
    /**
     * Sample code: Extensions_CreateOrUpdate.
     * 
     * @param manager Entry point to AgriFoodManager.
     */
    public static void extensionsCreateOrUpdate(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.extensions().define("provider.extension")
            .withExistingFarmBeat("examples-rg", "examples-farmbeatsResourceName").create();
    }
}
