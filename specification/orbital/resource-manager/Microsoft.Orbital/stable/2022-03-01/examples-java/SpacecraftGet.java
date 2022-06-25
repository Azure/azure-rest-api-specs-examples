import com.azure.core.util.Context;

/** Samples for Spacecrafts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftGet.json
     */
    /**
     * Sample code: Get Spacecraft.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void getSpacecraft(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.spacecrafts().getByResourceGroupWithResponse("contoso-Rgp", "CONTOSO_SAT", Context.NONE);
    }
}
