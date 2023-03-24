import com.azure.resourcemanager.digitaltwins.models.ConnectionProperties;
import com.azure.resourcemanager.digitaltwins.models.ConnectionPropertiesPrivateLinkServiceConnectionState;
import com.azure.resourcemanager.digitaltwins.models.PrivateEndpointConnection;
import com.azure.resourcemanager.digitaltwins.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/PrivateEndpointConnectionPut_example.json
     */
    /**
     * Sample code: Update the status of a private endpoint connection with the given name.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void updateTheStatusOfAPrivateEndpointConnectionWithTheGivenName(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        PrivateEndpointConnection resource =
            manager
                .privateEndpointConnections()
                .getWithResponse(
                    "resRg", "myDigitalTwinsService", "myPrivateConnection", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new ConnectionProperties()
                    .withPrivateLinkServiceConnectionState(
                        new ConnectionPropertiesPrivateLinkServiceConnectionState()
                            .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                            .withDescription("Approved by johndoe@company.com.")))
            .apply();
    }
}
