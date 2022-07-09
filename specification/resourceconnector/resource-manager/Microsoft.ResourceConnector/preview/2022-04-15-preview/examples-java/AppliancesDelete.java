import com.azure.core.util.Context;

/** Samples for Appliances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesDelete.json
     */
    /**
     * Sample code: Delete Appliance.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void deleteAppliance(com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().delete("testresourcegroup", "appliance01", Context.NONE);
    }
}
