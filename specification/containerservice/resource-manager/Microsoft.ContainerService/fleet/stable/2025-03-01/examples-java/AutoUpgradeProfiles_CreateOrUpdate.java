
import com.azure.resourcemanager.containerservicefleet.models.UpgradeChannel;

/**
 * Samples for AutoUpgradeProfiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/AutoUpgradeProfiles_CreateOrUpdate.json
     */
    /**
     * Sample code: Create an AutoUpgradeProfile.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void createAnAutoUpgradeProfile(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.autoUpgradeProfiles().define("autoupgradeprofile1").withExistingFleet("rg1", "fleet1")
            .withChannel(UpgradeChannel.STABLE).create();
    }
}
