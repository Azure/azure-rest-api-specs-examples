import com.azure.core.util.Context;

/** Samples for IotHubResource ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listkeys.json
     */
    /**
     * Sample code: IotHubResource_ListKeys.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceListKeys(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().listKeys("myResourceGroup", "testHub", Context.NONE);
    }
}
