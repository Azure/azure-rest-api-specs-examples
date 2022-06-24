import com.azure.core.util.Context;

/** Samples for Contacts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactGet.json
     */
    /**
     * Sample code: Get Contact.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void getContact(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contacts().getWithResponse("contoso-Rgp", "CONTOSO_SAT", "contact1", Context.NONE);
    }
}
