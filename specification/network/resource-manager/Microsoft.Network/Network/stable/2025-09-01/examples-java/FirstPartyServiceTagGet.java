
/**
 * Samples for FirstPartyServiceTags GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FirstPartyServiceTagGet.json
     */
    /**
     * Sample code: Get first party service tag.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getFirstPartyServiceTag(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirstPartyServiceTags().getByResourceGroupWithResponse("rg1", "myServiceTag",
            com.azure.core.util.Context.NONE);
    }
}
