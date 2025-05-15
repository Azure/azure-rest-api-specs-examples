
/**
 * Samples for AutoUpgradeProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/AutoUpgradeProfiles_Get.json
     */
    /**
     * Sample code: Gets an AutoUpgradeProfile resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getsAnAutoUpgradeProfileResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.autoUpgradeProfiles().getWithResponse("rg1", "fleet1", "autoupgradeprofile1",
            com.azure.core.util.Context.NONE);
    }
}
