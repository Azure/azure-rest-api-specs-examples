import com.azure.core.util.Context;

/** Samples for DigitalTwinsEndpoint Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/DigitalTwinsEndpointGet_WithIdentity_example.json
     */
    /**
     * Sample code: Get a DigitalTwinsInstance endpoint with identity.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getADigitalTwinsInstanceEndpointWithIdentity(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwinsEndpoints().getWithResponse("resRg", "myDigitalTwinsService", "myServiceBus", Context.NONE);
    }
}
