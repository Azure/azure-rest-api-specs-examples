/** Samples for PrivateLinkResources ListByConfigurationStore. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/PrivateLinkResourcesListByConfigurationStore.json
     */
    /**
     * Sample code: PrivateLinkResources_ListGroupIds.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void privateLinkResourcesListGroupIds(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .privateLinkResources()
            .listByConfigurationStore("myResourceGroup", "contoso", com.azure.core.util.Context.NONE);
    }
}
