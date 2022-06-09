```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.Function;
import com.azure.resourcemanager.streamanalytics.models.JavaScriptFunctionBinding;
import com.azure.resourcemanager.streamanalytics.models.ScalarFunctionProperties;

/** Samples for Functions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Update_JavaScript.json
     */
    /**
     * Sample code: Update a JavaScript function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAJavaScriptFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Function resource =
            manager.functions().getWithResponse("sjrg1637", "sj8653", "function8197", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new ScalarFunctionProperties()
                    .withBinding(new JavaScriptFunctionBinding().withScript("function (a, b) { return a * b; }")))
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
