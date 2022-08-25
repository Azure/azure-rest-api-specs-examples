import com.azure.core.util.Context;

/** Samples for CommunicationServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/communicationServices/listBySubscription.json
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
