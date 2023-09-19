/** Samples for IotHubResource GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_get.json
     */
    /**
     * Sample code: IotHubResource_Get.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGet(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .getByResourceGroupWithResponse("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
