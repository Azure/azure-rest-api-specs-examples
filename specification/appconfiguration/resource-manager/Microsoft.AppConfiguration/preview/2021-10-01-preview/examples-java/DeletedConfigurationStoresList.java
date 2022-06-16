import com.azure.core.util.Context;

/** Samples for ConfigurationStores ListDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/DeletedConfigurationStoresList.json
     */
    /**
     * Sample code: DeletedConfigurationStores_List.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void deletedConfigurationStoresList(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().listDeleted(Context.NONE);
    }
}
