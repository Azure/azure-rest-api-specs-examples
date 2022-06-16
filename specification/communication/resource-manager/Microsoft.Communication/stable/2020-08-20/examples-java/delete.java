import com.azure.core.util.Context;

/** Samples for CommunicationService Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/delete.json
     */
    /**
     * Sample code: Delete resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void deleteResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().delete("MyResourceGroup", "MyCommunicationResource", Context.NONE);
    }
}
