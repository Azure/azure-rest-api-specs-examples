
/**
 * Samples for AutoUpgradeProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/
     * examples/AutoUpgradeProfiles_Delete.json
     */
    /**
     * Sample code: Delete an AutoUpgradeProfile resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void deleteAnAutoUpgradeProfileResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.autoUpgradeProfiles().delete("rg1", "fleet1", "autoupgradeprofile1", null,
            com.azure.core.util.Context.NONE);
    }
}
