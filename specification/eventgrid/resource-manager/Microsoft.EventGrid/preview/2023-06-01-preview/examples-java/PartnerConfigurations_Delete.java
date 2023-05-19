/** Samples for PartnerConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerConfigurations_Delete.json
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
