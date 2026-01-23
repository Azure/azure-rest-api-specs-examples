
/**
 * Samples for EmailServices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/emailServices/
     * listByResourceGroup.json
     */
    /**
     * Sample code: List EmailService resources by resource group.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        listEmailServiceResourcesByResourceGroup(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.emailServices().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
