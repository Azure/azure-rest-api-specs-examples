
import com.azure.resourcemanager.communication.models.NameAvailabilityParameters;

/**
 * Samples for CommunicationServices CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/checkNameAvailabilityUnavailable.json
     */
    /**
     * Sample code: Check name availability unavailable.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        checkNameAvailabilityUnavailable(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().checkNameAvailabilityWithResponse(new NameAvailabilityParameters()
            .withName("MyCommunicationService").withType("Microsoft.Communication/CommunicationServices"),
            com.azure.core.util.Context.NONE);
    }
}
