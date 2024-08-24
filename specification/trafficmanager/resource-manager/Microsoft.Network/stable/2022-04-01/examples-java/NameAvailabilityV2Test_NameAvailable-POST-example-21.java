
import com.azure.resourcemanager.trafficmanager.models.CheckTrafficManagerRelativeDnsNameAvailabilityParameters;

/**
 * Samples for Profiles CheckTrafficManagerNameAvailabilityV2.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/
     * NameAvailabilityV2Test_NameAvailable-POST-example-21.json
     */
    /**
     * Sample code: NameAvailabilityV2Test_NameAvailablePOST21.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameAvailabilityV2TestNameAvailablePOST21(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getProfiles()
            .checkTrafficManagerNameAvailabilityV2WithResponse(
                new CheckTrafficManagerRelativeDnsNameAvailabilityParameters().withName("azsmnet5403")
                    .withType("microsoft.network/trafficmanagerprofiles"),
                com.azure.core.util.Context.NONE);
    }
}
