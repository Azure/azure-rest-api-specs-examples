import com.azure.core.util.Context;

/** Samples for CommunicationServices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/communicationServices/get.json
     */
    /**
     * Sample code: Get resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void getResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .getByResourceGroupWithResponse("MyResourceGroup", "MyCommunicationResource", Context.NONE);
    }
}
