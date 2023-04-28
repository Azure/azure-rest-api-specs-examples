/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/PrivateLinkResourceGet.json
     */
    /**
     * Sample code: PrivateLinkResources_Get.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void privateLinkResourcesGet(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .privateLinkResources()
            .getWithResponse("myResourceGroup", "contoso", "configurationStores", com.azure.core.util.Context.NONE);
    }
}
