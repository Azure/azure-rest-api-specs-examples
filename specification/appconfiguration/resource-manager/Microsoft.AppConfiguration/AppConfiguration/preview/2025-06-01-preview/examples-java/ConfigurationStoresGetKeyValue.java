
/**
 * Samples for KeyValues Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresGetKeyValue.json
     */
    /**
     * Sample code: KeyValues_Get.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.keyValues().getWithResponse("myResourceGroup", "contoso", "myKey$myLabel",
            com.azure.core.util.Context.NONE);
    }
}
