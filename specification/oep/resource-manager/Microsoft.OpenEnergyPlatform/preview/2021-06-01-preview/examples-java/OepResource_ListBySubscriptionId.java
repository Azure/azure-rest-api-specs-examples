import com.azure.core.util.Context;

/** Samples for EnergyServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_ListBySubscriptionId.json
     */
    /**
     * Sample code: OepResource_ListBySubscriptionId.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceListBySubscriptionId(com.azure.resourcemanager.oep.OepManager manager) {
        manager.energyServices().list(Context.NONE);
    }
}
