import com.azure.core.util.Context;

/** Samples for CommunicationService ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/listKeys.json
     */
    /**
     * Sample code: List keys.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void listKeys(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .listKeysWithResponse("MyResourceGroup", "MyCommunicationResource", Context.NONE);
    }
}
