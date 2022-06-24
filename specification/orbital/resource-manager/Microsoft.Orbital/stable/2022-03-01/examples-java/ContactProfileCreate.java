/** Samples for ContactProfiles CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileCreate.json
     */
    /**
     * Sample code: Create a contact profile.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void createAContactProfile(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager
            .contactProfiles()
            .define("CONTOSO-CP")
            .withRegion("eastus2")
            .withExistingResourceGroup("contoso-Rgp")
            .create();
    }
}
