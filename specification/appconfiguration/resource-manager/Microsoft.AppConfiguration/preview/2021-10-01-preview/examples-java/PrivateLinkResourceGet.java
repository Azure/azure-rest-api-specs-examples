import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/PrivateLinkResourceGet.json
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
            .getWithResponse("myResourceGroup", "contoso", "configurationStores", Context.NONE);
    }
}
