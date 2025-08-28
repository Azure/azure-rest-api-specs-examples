
import com.azure.resourcemanager.containerservicefleet.models.GatePatch;
import com.azure.resourcemanager.containerservicefleet.models.GatePatchProperties;
import com.azure.resourcemanager.containerservicefleet.models.GateState;

/**
 * Samples for Gates Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/Gates_Update.json
     */
    /**
     * Sample code: Updates a Gate resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        updatesAGateResource(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.gates().update("rg1", "fleet1", "12345678-910a-bcde-f000-000000000000",
            new GatePatch().withProperties(new GatePatchProperties().withState(GateState.COMPLETED)), null, null,
            com.azure.core.util.Context.NONE);
    }
}
