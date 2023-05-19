/** Samples for PartnerNamespaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerNamespaces_ListBySubscription.json
     */
    /**
     * Sample code: PartnerNamespaces_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerNamespacesListBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerNamespaces().list(null, null, com.azure.core.util.Context.NONE);
    }
}
