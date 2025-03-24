
/**
 * Samples for PartnerConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerConfigurations_Delete.json
     */
    /**
     * Sample code: PartnerConfigurations_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerConfigurationsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerConfigurations().delete("examplerg", com.azure.core.util.Context.NONE);
    }
}
