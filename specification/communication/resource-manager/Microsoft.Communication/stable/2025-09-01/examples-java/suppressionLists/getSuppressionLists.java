
/**
 * Samples for SuppressionLists ListByDomain.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/
     * getSuppressionLists.json
     */
    /**
     * Sample code: Get all SuppressionLists resources.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        getAllSuppressionListsResources(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.suppressionLists().listByDomain("contosoResourceGroup", "contosoEmailService", "contoso.com",
            com.azure.core.util.Context.NONE);
    }
}
