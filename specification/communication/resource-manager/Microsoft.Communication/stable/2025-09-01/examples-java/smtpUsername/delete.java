
/**
 * Samples for SmtpUsernames Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/smtpUsername/
     * delete.json
     */
    /**
     * Sample code: Delete a SmtpUsername resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        deleteASmtpUsernameResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.smtpUsernames().deleteWithResponse("MyResourceGroup", "contosoACSService", "smtpusername1",
            com.azure.core.util.Context.NONE);
    }
}
