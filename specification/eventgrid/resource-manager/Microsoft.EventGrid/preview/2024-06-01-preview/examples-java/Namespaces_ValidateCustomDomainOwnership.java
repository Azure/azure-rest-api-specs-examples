
/**
 * Samples for Namespaces ValidateCustomDomainOwnership.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * Namespaces_ValidateCustomDomainOwnership.json
     */
    /**
     * Sample code: Namespaces_ValidateCustomDomainOwnership.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        namespacesValidateCustomDomainOwnership(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaces().validateCustomDomainOwnership("examplerg", "exampleNamespaceName1",
            com.azure.core.util.Context.NONE);
    }
}
