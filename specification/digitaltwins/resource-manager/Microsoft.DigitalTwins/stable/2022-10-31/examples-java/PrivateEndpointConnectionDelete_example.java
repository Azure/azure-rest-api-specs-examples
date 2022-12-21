import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/PrivateEndpointConnectionDelete_example.json
     */
    /**
     * Sample code: Delete private endpoint connection with the specified name.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void deletePrivateEndpointConnectionWithTheSpecifiedName(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .privateEndpointConnections()
            .delete("resRg", "myDigitalTwinsService", "myPrivateConnection", Context.NONE);
    }
}
