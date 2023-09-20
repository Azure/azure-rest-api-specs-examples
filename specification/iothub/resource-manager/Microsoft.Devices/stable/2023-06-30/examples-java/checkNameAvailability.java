import com.azure.resourcemanager.iothub.models.OperationInputs;

/** Samples for IotHubResource CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/checkNameAvailability.json
     */
    /**
     * Sample code: IotHubResource_CheckNameAvailability.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceCheckNameAvailability(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .checkNameAvailabilityWithResponse(
                new OperationInputs().withName("test-request"), com.azure.core.util.Context.NONE);
    }
}
