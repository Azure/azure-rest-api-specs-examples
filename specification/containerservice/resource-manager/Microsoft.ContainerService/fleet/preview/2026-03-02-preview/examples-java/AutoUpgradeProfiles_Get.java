
/**
 * Samples for AutoUpgradeProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/AutoUpgradeProfiles_Get.json
     */
    /**
     * Sample code: Gets an AutoUpgradeProfile resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getsAnAutoUpgradeProfileResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.autoUpgradeProfiles().getWithResponse("rgfleets", "fleet1", "autoupgradeprofile1",
            com.azure.core.util.Context.NONE);
    }
}
