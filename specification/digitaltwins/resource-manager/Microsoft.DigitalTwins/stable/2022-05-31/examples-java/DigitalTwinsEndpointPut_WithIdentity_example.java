import com.azure.resourcemanager.digitaltwins.models.AuthenticationType;
import com.azure.resourcemanager.digitaltwins.models.ServiceBus;

/** Samples for DigitalTwinsEndpoint CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsEndpointPut_WithIdentity_example.json
     */
    /**
     * Sample code: Put a DigitalTwinsEndpoint resource with identity.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void putADigitalTwinsEndpointResourceWithIdentity(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwinsEndpoints()
            .define("myServiceBus")
            .withExistingDigitalTwinsInstance("resRg", "myDigitalTwinsService")
            .withProperties(
                new ServiceBus()
                    .withAuthenticationType(AuthenticationType.IDENTITY_BASED)
                    .withEndpointUri("sb://mysb.servicebus.windows.net/")
                    .withEntityPath("mysbtopic"))
            .create();
    }
}
