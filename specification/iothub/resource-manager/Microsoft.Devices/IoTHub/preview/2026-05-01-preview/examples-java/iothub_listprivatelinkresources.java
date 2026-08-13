
/**
 * Samples for PrivateLinkResourcesOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_listprivatelinkresources.json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void privateLinkResourcesList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.privateLinkResourcesOperations().listWithResponse("myResourceGroup", "testHub",
            com.azure.core.util.Context.NONE);
    }
}
