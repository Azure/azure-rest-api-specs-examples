
/**
 * Samples for PartnerConfigurations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerConfigurations_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerConfigurations_ListByResourceGroup.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        partnerConfigurationsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerConfigurations().listByResourceGroup("examplerg", com.azure.core.util.Context.NONE);
    }
}
