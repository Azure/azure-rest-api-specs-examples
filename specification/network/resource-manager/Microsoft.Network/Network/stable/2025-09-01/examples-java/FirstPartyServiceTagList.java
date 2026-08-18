
/**
 * Samples for FirstPartyServiceTags ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FirstPartyServiceTagList.json
     */
    /**
     * Sample code: List first party service tags in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listFirstPartyServiceTagsInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirstPartyServiceTags().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
