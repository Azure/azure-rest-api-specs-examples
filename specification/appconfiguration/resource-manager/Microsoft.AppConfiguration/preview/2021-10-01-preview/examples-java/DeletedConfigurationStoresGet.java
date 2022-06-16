import com.azure.core.util.Context;

/** Samples for ConfigurationStores GetDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/DeletedConfigurationStoresGet.json
     */
    /**
     * Sample code: DeletedConfigurationStores_Get.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void deletedConfigurationStoresGet(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().getDeletedWithResponse("westus", "contoso", Context.NONE);
    }
}
