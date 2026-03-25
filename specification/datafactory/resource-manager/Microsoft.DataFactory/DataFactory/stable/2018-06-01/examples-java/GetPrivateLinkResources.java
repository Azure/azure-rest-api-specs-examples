
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/GetPrivateLinkResources.json
     */
    /**
     * Sample code: Get private link resources of a site.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        getPrivateLinkResourcesOfASite(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.privateLinkResources().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
