/** Samples for PartnerConfigurations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerConfigurations_Get.json
     */
    /**
     * Sample code: PartnerConfigurations_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerConfigurationsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerConfigurations().getByResourceGroupWithResponse("examplerg", com.azure.core.util.Context.NONE);
    }
}
