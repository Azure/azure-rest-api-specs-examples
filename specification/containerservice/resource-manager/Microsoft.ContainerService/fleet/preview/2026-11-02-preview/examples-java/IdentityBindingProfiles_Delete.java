
/**
 * Samples for IdentityBindingProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-02-preview/IdentityBindingProfiles_Delete.json
     */
    /**
     * Sample code: Delete an identity binding profile.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void deleteAnIdentityBindingProfile(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.identityBindingProfiles().delete("rg1", "fleet1", "identitybinding1", "uktvayathbu",
            com.azure.core.util.Context.NONE);
    }
}
