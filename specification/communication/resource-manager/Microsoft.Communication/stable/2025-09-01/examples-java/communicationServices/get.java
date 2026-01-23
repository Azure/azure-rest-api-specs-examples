
/**
 * Samples for CommunicationServices GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/get.json
     */
    /**
     * Sample code: Get resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void getResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().getByResourceGroupWithResponse("MyResourceGroup", "MyCommunicationResource",
            com.azure.core.util.Context.NONE);
    }
}
