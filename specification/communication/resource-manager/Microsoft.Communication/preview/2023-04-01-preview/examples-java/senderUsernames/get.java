/** Samples for SenderUsernames Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/senderUsernames/get.json
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
            .getWithResponse(
                "contosoResourceGroup",
                "contosoEmailService",
                "contoso.com",
                "contosoNewsAlerts",
                com.azure.core.util.Context.NONE);
    }
}
