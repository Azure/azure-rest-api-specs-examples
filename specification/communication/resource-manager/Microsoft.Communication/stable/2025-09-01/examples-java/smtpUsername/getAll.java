
/**
 * Samples for SmtpUsernames List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/smtpUsername/
     * getAll.json
     */
    /**
     * Sample code: Get all SmtpUsername resources for a CommunicationService resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void getAllSmtpUsernameResourcesForACommunicationServiceResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.smtpUsernames().list("contosoResourceGroup", "contosoACSService", com.azure.core.util.Context.NONE);
    }
}
