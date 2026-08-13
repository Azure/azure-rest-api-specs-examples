
/**
 * Samples for PrivateLinkResourcesOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_getprivatelinkresources.json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void privateLinkResourcesList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.privateLinkResourcesOperations().getWithResponse("myResourceGroup", "testHub", "iotHub",
            com.azure.core.util.Context.NONE);
    }
}
