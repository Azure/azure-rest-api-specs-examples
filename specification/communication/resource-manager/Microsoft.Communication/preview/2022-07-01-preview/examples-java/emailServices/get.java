import com.azure.core.util.Context;

/** Samples for EmailServices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/emailServices/get.json
     */
    /**
     * Sample code: Get EmailService resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void getEmailServiceResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .emailServices()
            .getByResourceGroupWithResponse("MyResourceGroup", "MyEmailServiceResource", Context.NONE);
    }
}
