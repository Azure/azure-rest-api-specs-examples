import com.azure.core.util.Context;

/** Samples for CommunicationService ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/listByResourceGroup.json
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
