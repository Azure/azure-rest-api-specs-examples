import com.azure.core.util.Context;

/** Samples for Domains GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/Domains_Get.json
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
