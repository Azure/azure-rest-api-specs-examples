
/**
 * Samples for PartnerNamespaces ListSharedAccessKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerNamespaces_ListSharedAccessKeys.json
     */
    /**
     * Sample code: PartnerNamespaces_ListSharedAccessKeys.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        partnerNamespacesListSharedAccessKeys(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerNamespaces().listSharedAccessKeysWithResponse("examplerg", "examplePartnerNamespaceName1",
            com.azure.core.util.Context.NONE);
    }
}
