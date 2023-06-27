import com.azure.resourcemanager.trafficmanager.models.CheckTrafficManagerRelativeDnsNameAvailabilityParameters;

/** Samples for Profiles CheckTrafficManagerRelativeDnsNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/NameAvailabilityTest_NameAvailable-POST-example-21.json
     */
    /**
     * Sample code: NameAvailabilityTest_NameAvailablePOST21.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameAvailabilityTestNameAvailablePOST21(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .checkTrafficManagerRelativeDnsNameAvailabilityWithResponse(
                new CheckTrafficManagerRelativeDnsNameAvailabilityParameters()
                    .withName("azsmnet5403")
                    .withType("microsoft.network/trafficmanagerprofiles"),
                com.azure.core.util.Context.NONE);
    }
}
