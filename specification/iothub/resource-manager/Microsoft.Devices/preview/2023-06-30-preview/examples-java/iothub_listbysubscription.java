/** Samples for IotHubResource List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_listbysubscription.json
     */
    /**
     * Sample code: IotHubResource_ListBySubscription.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceListBySubscription(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().list(com.azure.core.util.Context.NONE);
    }
}
