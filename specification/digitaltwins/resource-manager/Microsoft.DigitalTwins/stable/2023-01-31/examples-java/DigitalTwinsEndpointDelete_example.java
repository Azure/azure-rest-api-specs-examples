
/**
 * Samples for DigitalTwinsEndpoint Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/
     * DigitalTwinsEndpointDelete_example.json
     */
    /**
     * Sample code: Delete a DigitalTwinsInstance endpoint.
     * 
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void
        deleteADigitalTwinsInstanceEndpoint(com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwinsEndpoints().delete("resRg", "myDigitalTwinsService", "myendpoint",
            com.azure.core.util.Context.NONE);
    }
}
