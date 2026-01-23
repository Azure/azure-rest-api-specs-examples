
/**
 * Samples for SenderUsernames CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/senderUsernames/
     * createOrUpdate.json
     */
    /**
     * Sample code: Create or update SenderUsernames resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        createOrUpdateSenderUsernamesResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.senderUsernames().define("contosoNewsAlerts")
            .withExistingDomain("contosoResourceGroup", "contosoEmailService", "contoso.com")
            .withUsername("contosoNewsAlerts").withDisplayName("Contoso News Alerts").create();
    }
}
