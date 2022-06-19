import com.azure.core.util.Context;

/** Samples for KeyValues Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresDeleteKeyValue.json
     */
    /**
     * Sample code: KeyValues_Delete.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesDelete(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.keyValues().delete("myResourceGroup", "contoso", "myKey$myLabel", Context.NONE);
    }
}
