import com.azure.core.util.Context;

/** Samples for Contacts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactDelete.json
     */
    /**
     * Sample code: Delete Contact.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void deleteContact(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contacts().delete("contoso-Rgp", "CONTOSO_SAT", "contact1", Context.NONE);
    }
}
