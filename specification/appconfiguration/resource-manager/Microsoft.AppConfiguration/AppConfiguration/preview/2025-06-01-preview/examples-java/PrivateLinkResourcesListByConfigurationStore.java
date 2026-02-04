
/**
 * Samples for PrivateLinkResources ListByConfigurationStore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/PrivateLinkResourcesListByConfigurationStore.json
     */
    /**
     * Sample code: PrivateLinkResources_ListGroupIds.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        privateLinkResourcesListGroupIds(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.privateLinkResources().listByConfigurationStore("myResourceGroup", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
