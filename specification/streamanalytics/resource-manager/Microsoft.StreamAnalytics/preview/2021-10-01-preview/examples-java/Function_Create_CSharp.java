
import com.azure.resourcemanager.streamanalytics.models.CSharpFunctionBinding;
import com.azure.resourcemanager.streamanalytics.models.FunctionInput;
import com.azure.resourcemanager.streamanalytics.models.FunctionOutput;
import com.azure.resourcemanager.streamanalytics.models.ScalarFunctionProperties;
import com.azure.resourcemanager.streamanalytics.models.UpdateMode;
import java.util.Arrays;

/**
 * Samples for Functions CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Create_CSharp.json
     */
    /**
     * Sample code: Create a CLRUdf function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createACLRUdfFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().define("function588").withExistingStreamingjob("sjrg", "sjName")
            .withProperties(new ScalarFunctionProperties()
                .withInputs(Arrays.asList(new FunctionInput().withDataType("nvarchar(max)")))
                .withOutput(new FunctionOutput().withDataType("nvarchar(max)"))
                .withBinding(new CSharpFunctionBinding().withDllPath("ASAEdgeApplication2_CodeBehind")
                    .withClassProperty("ASAEdgeUDFDemo.Class1").withMethod("SquareFunction")
                    .withUpdateMode(UpdateMode.STATIC)))
            .create();
    }
}
