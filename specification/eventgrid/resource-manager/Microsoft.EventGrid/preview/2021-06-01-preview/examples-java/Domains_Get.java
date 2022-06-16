import com.azure.core.util.Context;

/** Samples for Domains GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/Domains_Get.json
     */
    /**
     * Sample code: Domains_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domains().getByResourceGroupWithResponse("examplerg", "exampledomain2", Context.NONE);
    }
}
