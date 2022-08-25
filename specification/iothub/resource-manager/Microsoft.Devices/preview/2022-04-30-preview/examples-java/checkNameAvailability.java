import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.OperationInputs;

/** Samples for IotHubResource CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/checkNameAvailability.json
     */
    /**
     * Sample code: IotHubResource_CheckNameAvailability.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceCheckNameAvailability(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .checkNameAvailabilityWithResponse(new OperationInputs().withName("test-request"), Context.NONE);
    }
}
