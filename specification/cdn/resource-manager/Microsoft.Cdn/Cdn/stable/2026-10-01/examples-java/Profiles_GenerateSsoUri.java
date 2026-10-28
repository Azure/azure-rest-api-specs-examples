
/**
 * Samples for Profiles GenerateSsoUri.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_GenerateSsoUri.json
     */
    /**
     * Sample code: Profiles_GenerateSsoUri.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesGenerateSsoUri(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().generateSsoUriWithResponse("RG", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
