
/**
 * Samples for EmailServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/emailServices/
     * createOrUpdate.json
     */
    /**
     * Sample code: Create or update EmailService resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        createOrUpdateEmailServiceResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.emailServices().define("MyEmailServiceResource").withRegion("Global")
            .withExistingResourceGroup("MyResourceGroup").withDataLocation("United States").create();
    }
}
