
/**
 * Samples for IdentityBindingProfiles ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-02-preview/IdentityBindingProfiles_ListByFleet.json
     */
    /**
     * Sample code: List identity binding profiles by fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listIdentityBindingProfilesByFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.identityBindingProfiles().listByFleet("rg1", "fleet1", com.azure.core.util.Context.NONE);
    }
}
