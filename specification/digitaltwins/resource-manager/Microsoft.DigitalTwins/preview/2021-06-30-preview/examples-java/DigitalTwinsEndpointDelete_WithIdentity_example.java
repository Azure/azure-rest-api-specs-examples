import com.azure.core.util.Context;

/** Samples for DigitalTwinsEndpoint Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/DigitalTwinsEndpointDelete_WithIdentity_example.json
     */
    /**
     * Sample code: Delete a DigitalTwinsInstance endpoint with identity.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void deleteADigitalTwinsInstanceEndpointWithIdentity(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwinsEndpoints().delete("resRg", "myDigitalTwinsService", "myendpoint", Context.NONE);
    }
}
