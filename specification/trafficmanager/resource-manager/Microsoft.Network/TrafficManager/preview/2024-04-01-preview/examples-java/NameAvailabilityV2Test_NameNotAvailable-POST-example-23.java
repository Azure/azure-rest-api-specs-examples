
import com.azure.resourcemanager.trafficmanager.models.CheckTrafficManagerRelativeDnsNameAvailabilityParameters;

/**
 * Samples for Profiles CheckTrafficManagerNameAvailabilityV2.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/NameAvailabilityV2Test_NameNotAvailable-POST-example-23.json
     */
    /**
     * Sample code: NameAvailabilityV2Test_NameNotAvailablePOST23.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        nameAvailabilityV2TestNameNotAvailablePOST23(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().checkTrafficManagerNameAvailabilityV2WithResponse(
            new CheckTrafficManagerRelativeDnsNameAvailabilityParameters().withName("azsmnet4696")
                .withType("microsoft.network/trafficmanagerprofiles"),
            com.azure.core.util.Context.NONE);
    }
}
