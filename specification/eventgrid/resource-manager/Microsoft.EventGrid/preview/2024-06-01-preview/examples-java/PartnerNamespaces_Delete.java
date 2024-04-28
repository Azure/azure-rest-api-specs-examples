
/**
 * Samples for PartnerNamespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * PartnerNamespaces_Delete.json
     */
    /**
     * Sample code: PartnerNamespaces_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerNamespacesDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerNamespaces().delete("examplerg", "examplePartnerNamespaceName1",
            com.azure.core.util.Context.NONE);
    }
}
