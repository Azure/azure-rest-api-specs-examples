import com.azure.resourcemanager.voiceservices.models.CheckNameAvailabilityRequest;

/** Samples for NameAvailability CheckLocal. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/NameAvailability_CheckLocal.json
     */
    /**
     * Sample code: CheckLocalNameAvailability.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void checkLocalNameAvailability(
        com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        manager
            .nameAvailabilities()
            .checkLocalWithResponse(
                "useast",
                new CheckNameAvailabilityRequest()
                    .withName("myname")
                    .withType("Microsoft.VoiceServices/CommunicationsGateways"),
                com.azure.core.util.Context.NONE);
    }
}
