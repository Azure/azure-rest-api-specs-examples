import com.azure.core.util.Context;

/** Samples for CommunicationService List. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/listBySubscription.json
     */
    /**
     * Sample code: List by subscription.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void listBySubscription(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().list(Context.NONE);
    }
}
