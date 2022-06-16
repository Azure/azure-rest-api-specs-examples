import com.azure.core.util.Context;

/** Samples for EmailServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/emailServices/delete.json
     */
    /**
     * Sample code: Delete EmailService resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void deleteEmailServiceResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.emailServices().delete("MyResourceGroup", "MyEmailServiceResource", Context.NONE);
    }
}
