import com.azure.core.util.Context;

/** Samples for CommunicationServices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/communicationServices/listByResourceGroup.json
     */
    /**
     * Sample code: List by resource group.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void listByResourceGroup(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().listByResourceGroup("MyResourceGroup", Context.NONE);
    }
}
