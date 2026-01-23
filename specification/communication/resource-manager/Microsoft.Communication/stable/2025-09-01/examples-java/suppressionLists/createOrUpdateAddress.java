
/**
 * Samples for SuppressionListAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/
     * createOrUpdateAddress.json
     */
    /**
     * Sample code: CreateOrUpdate SuppressionListAddress resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateSuppressionListAddressResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.suppressionListAddresses().define("11112222-3333-4444-5555-aaaabbbbcccc")
            .withExistingSuppressionList("contosoResourceGroup", "contosoEmailService", "contoso.com",
                "aaaa1111-bbbb-2222-3333-aaaa11112222")
            .withEmail("newuser1@fabrikam.com").withFirstName("updatedFirstName").create();
    }
}
