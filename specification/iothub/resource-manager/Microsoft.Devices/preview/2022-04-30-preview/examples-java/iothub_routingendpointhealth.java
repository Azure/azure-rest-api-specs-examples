import com.azure.core.util.Context;

/** Samples for IotHubResource GetEndpointHealth. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/iothub_routingendpointhealth.json
     */
    /**
     * Sample code: IotHubResource_GetEndpointHealth.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetEndpointHealth(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getEndpointHealth("myResourceGroup", "testHub", Context.NONE);
    }
}
