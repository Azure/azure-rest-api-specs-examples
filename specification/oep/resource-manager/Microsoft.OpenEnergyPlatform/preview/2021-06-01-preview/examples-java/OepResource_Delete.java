import com.azure.core.util.Context;

/** Samples for EnergyServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_Delete.json
     */
    /**
     * Sample code: OepResource_Delete.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceDelete(com.azure.resourcemanager.oep.OepManager manager) {
        manager.energyServices().delete("DummyResourceGroupName", "DummyResourceName", Context.NONE);
    }
}
