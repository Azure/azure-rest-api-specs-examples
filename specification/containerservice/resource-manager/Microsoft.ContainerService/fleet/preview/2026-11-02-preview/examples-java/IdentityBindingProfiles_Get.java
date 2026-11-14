
/**
 * Samples for IdentityBindingProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-02-preview/IdentityBindingProfiles_Get.json
     */
    /**
     * Sample code: Get an identity binding profile.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getAnIdentityBindingProfile(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.identityBindingProfiles().getWithResponse("rg1", "fleet1", "identitybinding1",
            com.azure.core.util.Context.NONE);
    }
}
