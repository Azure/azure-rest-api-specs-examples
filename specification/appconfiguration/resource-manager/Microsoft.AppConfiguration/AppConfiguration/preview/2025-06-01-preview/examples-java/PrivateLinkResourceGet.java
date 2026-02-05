
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/PrivateLinkResourceGet.json
     */
    /**
     * Sample code: PrivateLinkResources_Get.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        privateLinkResourcesGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.privateLinkResources().getWithResponse("myResourceGroup", "contoso", "configurationStores",
            com.azure.core.util.Context.NONE);
    }
}
