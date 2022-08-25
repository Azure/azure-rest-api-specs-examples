import com.azure.core.util.Context;

/** Samples for Domains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/domains/delete.json
     */
    /**
     * Sample code: Delete Domains resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void deleteDomainsResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.domains().delete("MyResourceGroup", "MyEmailServiceResource", "mydomain.com", Context.NONE);
    }
}
