
import com.azure.resourcemanager.trafficmanager.models.CheckTrafficManagerRelativeDnsNameAvailabilityParameters;

/**
 * Samples for Profiles CheckTrafficManagerRelativeDnsNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/NameAvailabilityTest_NameNotAvailable-POST-example-23.json
     */
    /**
     * Sample code: NameAvailabilityTest_NameNotAvailablePOST23.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        nameAvailabilityTestNameNotAvailablePOST23(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().checkTrafficManagerRelativeDnsNameAvailabilityWithResponse(
            new CheckTrafficManagerRelativeDnsNameAvailabilityParameters().withName("azsmnet4696")
                .withType("microsoft.network/trafficmanagerprofiles"),
            com.azure.core.util.Context.NONE);
    }
}
