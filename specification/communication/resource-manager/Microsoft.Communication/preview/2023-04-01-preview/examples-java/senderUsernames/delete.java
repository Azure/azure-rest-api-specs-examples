/** Samples for SenderUsernames Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/senderUsernames/delete.json
     */
    /**
     * Sample code: Delete SenderUsernames resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void deleteSenderUsernamesResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .senderUsernames()
            .deleteWithResponse(
                "MyResourceGroup",
                "MyEmailServiceResource",
                "mydomain.com",
                "contosoNewsAlerts",
                com.azure.core.util.Context.NONE);
    }
}
