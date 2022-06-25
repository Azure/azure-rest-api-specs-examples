import com.azure.core.util.Context;

/** Samples for ContactProfiles List. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfilesBySubscriptionList.json
     */
    /**
     * Sample code: List of Contact Profiles.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void listOfContactProfiles(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contactProfiles().list("opaqueString", Context.NONE);
    }
}
