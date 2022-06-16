import com.azure.resourcemanager.streamanalytics.models.FunctionInput;
import com.azure.resourcemanager.streamanalytics.models.FunctionOutput;
import com.azure.resourcemanager.streamanalytics.models.JavaScriptFunctionBinding;
import com.azure.resourcemanager.streamanalytics.models.ScalarFunctionProperties;
import java.util.Arrays;

/** Samples for Functions CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Create_JavaScript.json
     */
    /**
     * Sample code: Create a JavaScript function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAJavaScriptFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .functions()
            .define("function8197")
            .withExistingStreamingjob("sjrg1637", "sj8653")
            .withProperties(
                new ScalarFunctionProperties()
                    .withInputs(Arrays.asList(new FunctionInput().withDataType("Any")))
                    .withOutput(new FunctionOutput().withDataType("Any"))
                    .withBinding(new JavaScriptFunctionBinding().withScript("function (x, y) { return x + y; }")))
            .create();
    }
}
