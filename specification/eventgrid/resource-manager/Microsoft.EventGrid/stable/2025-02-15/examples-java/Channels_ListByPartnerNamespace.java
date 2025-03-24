
/**
 * Samples for Channels ListByPartnerNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * Channels_ListByPartnerNamespace.json
     */
    /**
     * Sample code: Channels_ListByPartnerNamespace.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsListByPartnerNamespace(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.channels().listByPartnerNamespace("examplerg", "examplePartnerNamespaceName1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
