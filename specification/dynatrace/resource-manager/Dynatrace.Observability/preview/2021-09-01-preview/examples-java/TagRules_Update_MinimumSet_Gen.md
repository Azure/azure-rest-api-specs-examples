```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.TagRule;

/** Samples for TagRules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Update_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesUpdateMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        TagRule resource =
            manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE).getValue();
        resource.update().apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dynatrace_1.0.0-beta.1/sdk/dynatrace/azure-resourcemanager-dynatrace/README.md) on how to add the SDK to your project and authenticate.
