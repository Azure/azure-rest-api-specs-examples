import com.azure.core.util.Context;

/** Samples for DigitalTwinsEndpoint Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/DigitalTwinsEndpointGet_example.json
     */
    /**
     * Sample code: Get a DigitalTwinsInstance endpoint.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getADigitalTwinsInstanceEndpoint(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwinsEndpoints().getWithResponse("resRg", "myDigitalTwinsService", "myServiceBus", Context.NONE);
    }
}
