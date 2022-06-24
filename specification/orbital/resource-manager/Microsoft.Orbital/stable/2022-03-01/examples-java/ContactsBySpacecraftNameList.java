import com.azure.core.util.Context;

/** Samples for Contacts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactsBySpacecraftNameList.json
     */
    /**
     * Sample code: List of Spacecraft.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void listOfSpacecraft(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contacts().list("contoso-Rgp", "CONTOSO_SAT", "opaqueString", Context.NONE);
    }
}
