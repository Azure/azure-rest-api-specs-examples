/** Samples for Spacecrafts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/SpacecraftDelete.json
     */
    /**
     * Sample code: Delete Spacecraft.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void deleteSpacecraft(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.spacecrafts().delete("contoso-Rgp", "CONTOSO_SAT", com.azure.core.util.Context.NONE);
    }
}
