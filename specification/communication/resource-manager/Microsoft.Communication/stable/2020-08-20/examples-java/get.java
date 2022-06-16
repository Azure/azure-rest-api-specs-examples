import com.azure.core.util.Context;

/** Samples for CommunicationService GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/get.json
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
