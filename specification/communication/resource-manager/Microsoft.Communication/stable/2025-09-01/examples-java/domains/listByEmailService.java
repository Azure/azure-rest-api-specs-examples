
/**
 * Samples for Domains ListByEmailServiceResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/domains/
     * listByEmailService.json
     */
    /**
     * Sample code: List Domains resources by EmailServiceName.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void
        listDomainsResourcesByEmailServiceName(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.domains().listByEmailServiceResource("MyResourceGroup", "MyEmailServiceResource",
            com.azure.core.util.Context.NONE);
    }
}
