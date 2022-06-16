import com.azure.core.util.Context;

/** Samples for PartnerConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerConfigurations_ListBySubscription.json
     */
    /**
     * Sample code: PartnerConfigurations_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerConfigurationsListBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerConfigurations().list(null, null, Context.NONE);
    }
}
