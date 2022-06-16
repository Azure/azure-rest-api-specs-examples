import com.azure.core.util.Context;

/** Samples for PartnerConfigurations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerConfigurations_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerConfigurations_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerConfigurationsListByResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerConfigurations().listByResourceGroup("examplerg", Context.NONE);
    }
}
