import com.azure.core.util.Context;

/** Samples for KeyValues Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresGetKeyValue.json
     */
    /**
     * Sample code: KeyValues_Get.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.keyValues().getWithResponse("myResourceGroup", "contoso", "myKey$myLabel", Context.NONE);
    }
}
