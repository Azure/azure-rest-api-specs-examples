import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/PrivateEndpointConnectionByConnectionName_example.json
     */
    /**
     * Sample code: Get private endpoint connection properties for the given private endpoint.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getPrivateEndpointConnectionPropertiesForTheGivenPrivateEndpoint(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("resRg", "myDigitalTwinsService", "myPrivateConnection", Context.NONE);
    }
}
