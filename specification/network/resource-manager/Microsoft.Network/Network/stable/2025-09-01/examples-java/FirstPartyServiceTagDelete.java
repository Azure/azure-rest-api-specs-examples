
/**
 * Samples for FirstPartyServiceTags Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FirstPartyServiceTagDelete.json
     */
    /**
     * Sample code: Delete first party service tag.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteFirstPartyServiceTag(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirstPartyServiceTags().delete("rg1", "myServiceTag",
            com.azure.core.util.Context.NONE);
    }
}
