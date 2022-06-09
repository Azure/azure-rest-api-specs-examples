```java
import com.azure.resourcemanager.apimanagement.models.LoggerType;
import java.util.HashMap;
import java.util.Map;

/** Samples for Logger CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateEHLogger.json
     */
    /**
     * Sample code: ApiManagementCreateEHLogger.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateEHLogger(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .loggers()
            .define("eh1")
            .withExistingService("rg1", "apimService1")
            .withLoggerType(LoggerType.AZURE_EVENT_HUB)
            .withDescription("adding a new logger")
            .withCredentials(
                mapOf(
                    "name",
                    "hydraeventhub",
                    "connectionString",
                    "Endpoint=sb://hydraeventhub-ns.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=********="))
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
