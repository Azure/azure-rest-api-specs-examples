
import com.azure.resourcemanager.trafficmanager.models.CheckTrafficManagerRelativeDnsNameAvailabilityParameters;

/**
 * Samples for Profiles CheckTrafficManagerRelativeDnsNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/NameAvailabilityTest_NameAvailable-POST-example-21.json
     */
    /**
     * Sample code: NameAvailabilityTest_NameAvailablePOST21.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        nameAvailabilityTestNameAvailablePOST21(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().checkTrafficManagerRelativeDnsNameAvailabilityWithResponse(
            new CheckTrafficManagerRelativeDnsNameAvailabilityParameters().withName("azsmnet5403")
                .withType("microsoft.network/trafficmanagerprofiles"),
            com.azure.core.util.Context.NONE);
    }
}
