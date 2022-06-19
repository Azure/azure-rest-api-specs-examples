import com.azure.core.util.Context;

/** Samples for ConfigurationStores PurgeDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/DeletedConfigurationStoresPurge.json
     */
    /**
     * Sample code: Purge a deleted configuration store.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void purgeADeletedConfigurationStore(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().purgeDeleted("westus", "contoso", Context.NONE);
    }
}
