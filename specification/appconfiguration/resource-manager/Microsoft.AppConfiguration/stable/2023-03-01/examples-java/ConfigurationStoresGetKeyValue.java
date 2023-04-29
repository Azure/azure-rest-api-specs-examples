/** Samples for KeyValues Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresGetKeyValue.json
     */
    /**
     * Sample code: KeyValues_Get.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .keyValues()
            .getWithResponse("myResourceGroup", "contoso", "myKey$myLabel", com.azure.core.util.Context.NONE);
    }
}
