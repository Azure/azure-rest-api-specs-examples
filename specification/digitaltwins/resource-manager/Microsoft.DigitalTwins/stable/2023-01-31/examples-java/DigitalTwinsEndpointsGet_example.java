
/**
 * Samples for DigitalTwinsEndpoint List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/
     * DigitalTwinsEndpointsGet_example.json
     */
    /**
     * Sample code: Get a DigitalTwinsInstance endpoints.
     * 
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void
        getADigitalTwinsInstanceEndpoints(com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwinsEndpoints().list("resRg", "myDigitalTwinsService", com.azure.core.util.Context.NONE);
    }
}
