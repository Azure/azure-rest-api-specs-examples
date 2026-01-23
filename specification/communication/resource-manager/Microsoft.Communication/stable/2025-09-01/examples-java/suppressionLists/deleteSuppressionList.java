
/**
 * Samples for SuppressionLists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/
     * deleteSuppressionList.json
     */
    /**
     * Sample code: Delete a SuppressionLists resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        deleteASuppressionListsResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.suppressionLists().deleteWithResponse("MyResourceGroup", "MyEmailServiceResource", "mydomain.com",
            "aaaa1111-bbbb-2222-3333-aaaa11112222", com.azure.core.util.Context.NONE);
    }
}
