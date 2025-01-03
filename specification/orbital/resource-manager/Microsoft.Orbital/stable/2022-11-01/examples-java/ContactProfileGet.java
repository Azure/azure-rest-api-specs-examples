
/**
 * Samples for ContactProfiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactProfileGet.json
     */
    /**
     * Sample code: Get a contact profile.
     * 
     * @param manager Entry point to OrbitalManager.
     */
    public static void getAContactProfile(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contactProfiles().getByResourceGroupWithResponse("contoso-Rgp", "CONTOSO-CP",
            com.azure.core.util.Context.NONE);
    }
}
