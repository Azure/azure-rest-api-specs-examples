
/**
 * Samples for KeyValues Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresDeleteKeyValue.json
     */
    /**
     * Sample code: KeyValues_Delete.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesDelete(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.keyValues().delete("myResourceGroup", "contoso", "myKey$myLabel", com.azure.core.util.Context.NONE);
    }
}
