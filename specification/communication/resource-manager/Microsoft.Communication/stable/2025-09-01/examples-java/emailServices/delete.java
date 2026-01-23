
/**
 * Samples for EmailServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/emailServices/
     * delete.json
     */
    /**
     * Sample code: Delete EmailService resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        deleteEmailServiceResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.emailServices().delete("MyResourceGroup", "MyEmailServiceResource", com.azure.core.util.Context.NONE);
    }
}
