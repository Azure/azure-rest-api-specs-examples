/** Samples for PrivateLinkResourcesOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_listprivatelinkresources.json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void privateLinkResourcesList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .privateLinkResourcesOperations()
            .listWithResponse("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
