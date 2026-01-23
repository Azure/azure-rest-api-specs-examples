
/**
 * Samples for SuppressionLists Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/
     * getSuppressionList.json
     */
    /**
     * Sample code: Get a SuppressionList resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        getASuppressionListResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.suppressionLists().getWithResponse("contosoResourceGroup", "contosoEmailService", "contoso.com",
            "aaaa1111-bbbb-2222-3333-aaaa11112222", com.azure.core.util.Context.NONE);
    }
}
