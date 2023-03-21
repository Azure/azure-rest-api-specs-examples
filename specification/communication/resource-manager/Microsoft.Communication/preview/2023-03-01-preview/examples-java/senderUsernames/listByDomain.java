/** Samples for SenderUsernames ListByDomains. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-03-01-preview/examples/senderUsernames/listByDomain.json
     */
    /**
     * Sample code: Get SenderUsernames resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void getSenderUsernamesResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .senderUsernames()
            .listByDomains(
                "contosoResourceGroup", "contosoEmailService", "contoso.com", com.azure.core.util.Context.NONE);
    }
}
