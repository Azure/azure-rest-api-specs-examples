import com.azure.resourcemanager.communication.models.DomainManagement;

/** Samples for Domains CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/domains/createOrUpdate.json
     */
    /**
     * Sample code: Create or update Domains resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateDomainsResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .domains()
            .define("mydomain.com")
            .withRegion("Global")
            .withExistingEmailService("MyResourceGroup", "MyEmailServiceResource")
            .withDomainManagement(DomainManagement.CUSTOMER_MANAGED)
            .create();
    }
}
