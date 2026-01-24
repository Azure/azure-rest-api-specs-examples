
/**
 * Samples for SmtpUsernames CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/smtpUsername/
     * createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate SmtpUsername resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        createOrUpdateSmtpUsernameResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.smtpUsernames().define("smtpusername1")
            .withExistingCommunicationService("contosoResourceGroup", "contosoACSService")
            .withUsername("newuser1@contoso.com").withEntraApplicationId("aaaa1111-bbbb-2222-3333-aaaa111122bb")
            .withTenantId("aaaa1111-bbbb-2222-3333-aaaa11112222").create();
    }
}
