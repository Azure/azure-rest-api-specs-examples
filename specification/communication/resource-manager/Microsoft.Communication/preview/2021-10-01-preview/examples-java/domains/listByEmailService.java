import com.azure.core.util.Context;

/** Samples for Domains ListByEmailServiceResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/domains/listByEmailService.json
     */
    /**
     * Sample code: List Domains resources by EmailServiceName.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void listDomainsResourcesByEmailServiceName(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.domains().listByEmailServiceResource("MyResourceGroup", "MyEmailServiceResource", Context.NONE);
    }
}
