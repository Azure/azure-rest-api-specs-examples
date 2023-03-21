/** Samples for EmailServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-03-01-preview/examples/emailServices/listBySubscription.json
     */
    /**
     * Sample code: List EmailService resources by subscription.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void listEmailServiceResourcesBySubscription(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.emailServices().list(com.azure.core.util.Context.NONE);
    }
}
