/** Samples for IotHubResource Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_delete.json
     */
    /**
     * Sample code: IotHubResource_Delete.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceDelete(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().delete("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
