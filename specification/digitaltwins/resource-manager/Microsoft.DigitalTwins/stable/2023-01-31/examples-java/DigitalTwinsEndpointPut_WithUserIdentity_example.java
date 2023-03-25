import com.azure.resourcemanager.digitaltwins.models.AuthenticationType;
import com.azure.resourcemanager.digitaltwins.models.IdentityType;
import com.azure.resourcemanager.digitaltwins.models.ManagedIdentityReference;
import com.azure.resourcemanager.digitaltwins.models.ServiceBus;

/** Samples for DigitalTwinsEndpoint CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsEndpointPut_WithUserIdentity_example.json
     */
    /**
     * Sample code: Put a DigitalTwinsEndpoint resource with user assigned identity.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void putADigitalTwinsEndpointResourceWithUserAssignedIdentity(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwinsEndpoints()
            .define("myServiceBus")
            .withExistingDigitalTwinsInstance("resRg", "myDigitalTwinsService")
            .withProperties(
                new ServiceBus()
                    .withAuthenticationType(AuthenticationType.IDENTITY_BASED)
                    .withIdentity(
                        new ManagedIdentityReference()
                            .withType(IdentityType.USER_ASSIGNED)
                            .withUserAssignedIdentity(
                                "/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity"))
                    .withEndpointUri("sb://mysb.servicebus.windows.net/")
                    .withEntityPath("mysbtopic"))
            .create();
    }
}
