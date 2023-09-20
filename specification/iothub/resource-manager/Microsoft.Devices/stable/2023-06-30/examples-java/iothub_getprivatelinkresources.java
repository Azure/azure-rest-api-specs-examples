/** Samples for PrivateLinkResourcesOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_getprivatelinkresources.json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void privateLinkResourcesList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .privateLinkResourcesOperations()
            .getWithResponse("myResourceGroup", "testHub", "iotHub", com.azure.core.util.Context.NONE);
    }
}
