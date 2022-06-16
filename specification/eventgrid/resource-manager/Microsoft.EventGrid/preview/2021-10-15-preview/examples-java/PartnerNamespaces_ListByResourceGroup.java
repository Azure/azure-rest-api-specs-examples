import com.azure.core.util.Context;

/** Samples for PartnerNamespaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerNamespaces_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerNamespaces_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerNamespacesListByResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerNamespaces().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
