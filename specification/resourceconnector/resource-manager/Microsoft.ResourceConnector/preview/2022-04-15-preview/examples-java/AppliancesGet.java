import com.azure.core.util.Context;

/** Samples for Appliances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesGet.json
     */
    /**
     * Sample code: Get Appliance.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void getAppliance(com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().getByResourceGroupWithResponse("testresourcegroup", "appliance01", Context.NONE);
    }
}
