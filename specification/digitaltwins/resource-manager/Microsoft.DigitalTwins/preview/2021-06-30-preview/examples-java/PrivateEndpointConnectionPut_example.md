```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.digitaltwins.models.ConnectionProperties;
import com.azure.resourcemanager.digitaltwins.models.ConnectionPropertiesPrivateLinkServiceConnectionState;
import com.azure.resourcemanager.digitaltwins.models.PrivateEndpointConnection;
import com.azure.resourcemanager.digitaltwins.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/PrivateEndpointConnectionPut_example.json
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
                .getWithResponse("resRg", "myDigitalTwinsService", "myPrivateConnection", Context.NONE)
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-digitaltwins_1.0.0-beta.2/sdk/digitaltwins/azure-resourcemanager-digitaltwins/README.md) on how to add the SDK to your project and authenticate.
