import com.azure.resourcemanager.iothub.models.FailoverInput;

/** Samples for IotHub ManualFailover. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/IotHub_ManualFailover.json
     */
    /**
     * Sample code: IotHub_ManualFailover.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubManualFailover(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubs()
            .manualFailover(
                "testHub",
                "myResourceGroup",
                new FailoverInput().withFailoverRegion("testHub"),
                com.azure.core.util.Context.NONE);
    }
}
