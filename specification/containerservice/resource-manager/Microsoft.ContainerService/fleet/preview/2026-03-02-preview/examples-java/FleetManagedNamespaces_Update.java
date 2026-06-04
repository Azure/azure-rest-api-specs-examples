
import com.azure.resourcemanager.containerservicefleet.models.FleetManagedNamespace;

/**
 * Samples for FleetManagedNamespaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/FleetManagedNamespaces_Update.json
     */
    /**
     * Sample code: FleetManagedNamespaces_Update.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void fleetManagedNamespacesUpdate(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        FleetManagedNamespace resource = manager.fleetManagedNamespaces()
            .getWithResponse("rgfleets", "fleet1", "namespace1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
