
/**
 * Samples for SmtpUsernames Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/smtpUsername/get.
     * json
     */
    /**
     * Sample code: Get a SmtpUsername resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void getASmtpUsernameResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.smtpUsernames().getWithResponse("contosoResourceGroup", "contosoACSService", "smtpusername1",
            com.azure.core.util.Context.NONE);
    }
}
