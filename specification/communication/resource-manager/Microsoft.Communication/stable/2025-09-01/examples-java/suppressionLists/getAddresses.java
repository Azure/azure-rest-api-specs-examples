
/**
 * Samples for SuppressionListAddresses List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/
     * getAddresses.json
     */
    /**
     * Sample code: Get all SuppressionListAddresses resources for a SuppressionList resource.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void getAllSuppressionListAddressesResourcesForASuppressionListResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.suppressionListAddresses().list("contosoResourceGroup", "contosoEmailService", "contoso.com",
            "aaaa1111-bbbb-2222-3333-aaaa11112222", com.azure.core.util.Context.NONE);
    }
}
