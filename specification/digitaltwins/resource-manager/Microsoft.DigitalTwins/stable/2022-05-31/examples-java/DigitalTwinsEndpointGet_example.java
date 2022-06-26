import com.azure.core.util.Context;

/** Samples for DigitalTwinsEndpoint Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsEndpointGet_example.json
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
