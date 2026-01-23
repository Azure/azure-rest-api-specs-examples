
/**
 * Samples for CommunicationServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/delete.json
     */
    /**
     * Sample code: Delete resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void deleteResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().delete("MyResourceGroup", "MyCommunicationResource",
            com.azure.core.util.Context.NONE);
    }
}
