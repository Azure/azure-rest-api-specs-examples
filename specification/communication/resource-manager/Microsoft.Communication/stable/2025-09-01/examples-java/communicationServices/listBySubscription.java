
/**
 * Samples for CommunicationServices List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/listBySubscription.json
     */
    /**
     * Sample code: List by subscription.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void listBySubscription(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().list(com.azure.core.util.Context.NONE);
    }
}
