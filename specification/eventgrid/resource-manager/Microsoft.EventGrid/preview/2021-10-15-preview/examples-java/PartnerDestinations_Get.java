import com.azure.core.util.Context;

/** Samples for PartnerDestinations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerDestinations_Get.json
     */
    /**
     * Sample code: PartnerDestinations_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerDestinationsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerDestinations()
            .getByResourceGroupWithResponse("examplerg", "examplePartnerDestinationName1", Context.NONE);
    }
}
