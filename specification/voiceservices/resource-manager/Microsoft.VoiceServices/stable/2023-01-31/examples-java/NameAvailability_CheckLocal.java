import com.azure.resourcemanager.voiceservices.models.CheckNameAvailabilityRequest;

/** Samples for NameAvailability CheckLocal. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/NameAvailability_CheckLocal.json
     */
    /**
     * Sample code: CheckLocalNameAvailability.
     *
     * @param manager Entry point to VoiceservicesManager.
     */
    public static void checkLocalNameAvailability(
        com.azure.resourcemanager.voiceservices.VoiceservicesManager manager) {
        manager
            .nameAvailabilities()
            .checkLocalWithResponse(
                "useast",
                new CheckNameAvailabilityRequest()
                    .withName("myname")
                    .withType("Microsoft.VoiceServices/CommunicationsGateway"),
                com.azure.core.util.Context.NONE);
    }
}
