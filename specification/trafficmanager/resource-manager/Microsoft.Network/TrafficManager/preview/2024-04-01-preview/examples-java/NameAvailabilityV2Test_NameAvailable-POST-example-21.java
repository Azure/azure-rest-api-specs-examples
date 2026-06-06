
import com.azure.resourcemanager.trafficmanager.models.CheckTrafficManagerRelativeDnsNameAvailabilityParameters;

/**
 * Samples for Profiles CheckTrafficManagerNameAvailabilityV2.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/NameAvailabilityV2Test_NameAvailable-POST-example-21.json
     */
    /**
     * Sample code: NameAvailabilityV2Test_NameAvailablePOST21.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        nameAvailabilityV2TestNameAvailablePOST21(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().checkTrafficManagerNameAvailabilityV2WithResponse(
            new CheckTrafficManagerRelativeDnsNameAvailabilityParameters().withName("azsmnet5403")
                .withType("microsoft.network/trafficmanagerprofiles"),
            com.azure.core.util.Context.NONE);
    }
}
