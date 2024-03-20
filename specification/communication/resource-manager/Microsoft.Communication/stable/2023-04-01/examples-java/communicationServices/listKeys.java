
/**
 * Samples for CommunicationServices ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2023-04-01/examples/
     * communicationServices/listKeys.json
     */
    /**
     * Sample code: List keys.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void listKeys(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().listKeysWithResponse("MyResourceGroup", "MyCommunicationResource",
            com.azure.core.util.Context.NONE);
    }
}
