import com.azure.core.util.Context;
import com.azure.resourcemanager.oep.models.EnergyService;

/** Samples for EnergyServices Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_Update.json
     */
    /**
     * Sample code: OepResource_Update.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceUpdate(com.azure.resourcemanager.oep.OepManager manager) {
        EnergyService resource =
            manager
                .energyServices()
                .getByResourceGroupWithResponse("DummyResourceGroupName", "DummyResourceName", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
