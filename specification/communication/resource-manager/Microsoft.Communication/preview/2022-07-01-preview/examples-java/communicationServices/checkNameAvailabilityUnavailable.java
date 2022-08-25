import com.azure.core.util.Context;
import com.azure.resourcemanager.communication.models.NameAvailabilityParameters;

/** Samples for CommunicationServices CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/communicationServices/checkNameAvailabilityUnavailable.json
     */
    /**
     * Sample code: Check name availability unavailable.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void checkNameAvailabilityUnavailable(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .checkNameAvailabilityWithResponse(
                new NameAvailabilityParameters()
                    .withName("MyCommunicationService")
                    .withType("Microsoft.Communication/CommunicationServices"),
                Context.NONE);
    }
}
