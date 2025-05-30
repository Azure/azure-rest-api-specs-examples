
/**
 * Samples for Spacecrafts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/
     * SpacecraftsByResourceGroupList.json
     */
    /**
     * Sample code: List of Spacecraft by Resource Group.
     * 
     * @param manager Entry point to OrbitalManager.
     */
    public static void listOfSpacecraftByResourceGroup(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.spacecrafts().listByResourceGroup("contoso-Rgp", "opaqueString", com.azure.core.util.Context.NONE);
    }
}
