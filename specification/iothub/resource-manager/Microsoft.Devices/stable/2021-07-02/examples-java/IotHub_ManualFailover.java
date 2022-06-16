import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.FailoverInput;

/** Samples for IotHub ManualFailover. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/IotHub_ManualFailover.json
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
                "testHub", "myResourceGroup", new FailoverInput().withFailoverRegion("testHub"), Context.NONE);
    }
}
