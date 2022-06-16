import com.azure.core.util.Context;

/** Samples for KeyValues ListByConfigurationStore. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresListKeyValues.json
     */
    /**
     * Sample code: KeyValues_ListByConfigurationStore.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesListByConfigurationStore(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.keyValues().listByConfigurationStore("myResourceGroup", "contoso", null, Context.NONE);
    }
}
