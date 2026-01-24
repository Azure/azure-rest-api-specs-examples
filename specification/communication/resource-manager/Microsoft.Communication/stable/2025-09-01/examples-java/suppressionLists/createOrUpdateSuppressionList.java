
/**
 * Samples for SuppressionLists CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/
     * createOrUpdateSuppressionList.json
     */
    /**
     * Sample code: CreateOrUpdate SuppressionLists resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        createOrUpdateSuppressionListsResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.suppressionLists().define("aaaa1111-bbbb-2222-3333-aaaa11112222")
            .withExistingDomain("contosoResourceGroup", "contosoEmailService", "contoso.com")
            .withListName("contosoNewsAlerts").create();
    }
}
