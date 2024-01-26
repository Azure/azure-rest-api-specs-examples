
import com.azure.resourcemanager.cdn.models.CheckHostnameAvailabilityInput;

/** Samples for AfdProfiles CheckHostnameAvailability. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDProfiles_CheckHostNameAvailability
     * .json
     */
    /**
     * Sample code: AFDProfiles_CheckHostNameAvailability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDProfilesCheckHostNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdProfiles().checkHostnameAvailabilityWithResponse("RG",
            "profile1", new CheckHostnameAvailabilityInput().withHostname("www.someDomain.net"),
            com.azure.core.util.Context.NONE);
    }
}
