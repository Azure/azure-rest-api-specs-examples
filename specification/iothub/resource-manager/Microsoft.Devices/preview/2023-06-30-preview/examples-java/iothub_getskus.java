/** Samples for IotHubResource GetValidSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_getskus.json
     */
    /**
     * Sample code: IotHubResource_GetValidSkus.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetValidSkus(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getValidSkus("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
