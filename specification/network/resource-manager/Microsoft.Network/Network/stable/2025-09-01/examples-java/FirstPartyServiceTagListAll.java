
/**
 * Samples for FirstPartyServiceTags List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FirstPartyServiceTagListAll.json
     */
    /**
     * Sample code: List all first party service tags.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllFirstPartyServiceTags(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirstPartyServiceTags().list(com.azure.core.util.Context.NONE);
    }
}
