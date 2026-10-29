
import com.azure.resourcemanager.cdn.models.CheckHostnameAvailabilityInput;

/**
 * Samples for AfdProfiles CheckHostnameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDProfiles_CheckHostNameAvailability.json
     */
    /**
     * Sample code: AFDProfiles_CheckHostNameAvailability.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDProfilesCheckHostNameAvailability(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdProfiles().checkHostnameAvailabilityWithResponse("RG", "profile1",
            new CheckHostnameAvailabilityInput().withHostname("www.someDomain.net"), com.azure.core.util.Context.NONE);
    }
}
