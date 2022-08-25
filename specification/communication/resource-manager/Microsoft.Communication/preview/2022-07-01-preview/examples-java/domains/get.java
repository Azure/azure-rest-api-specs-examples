import com.azure.core.util.Context;

/** Samples for Domains Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/domains/get.json
     */
    /**
     * Sample code: Get Domains resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void getDomainsResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.domains().getWithResponse("MyResourceGroup", "MyEmailServiceResource", "mydomain.com", Context.NONE);
    }
}
