import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/PrivateEndpointConnectionsList_example.json
     */
    /**
     * Sample code: List private endpoint connection properties.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void listPrivateEndpointConnectionProperties(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.privateEndpointConnections().listWithResponse("resRg", "myDigitalTwinsService", Context.NONE);
    }
}
