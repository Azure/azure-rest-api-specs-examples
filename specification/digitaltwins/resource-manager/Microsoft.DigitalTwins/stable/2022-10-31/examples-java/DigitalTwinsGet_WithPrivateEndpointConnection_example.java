import com.azure.core.util.Context;

/** Samples for DigitalTwins GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/DigitalTwinsGet_WithPrivateEndpointConnection_example.json
     */
    /**
     * Sample code: Get a DigitalTwinsInstance resource with a private endpoint connection.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getADigitalTwinsInstanceResourceWithAPrivateEndpointConnection(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwins().getByResourceGroupWithResponse("resRg", "myDigitalTwinsService", Context.NONE);
    }
}
