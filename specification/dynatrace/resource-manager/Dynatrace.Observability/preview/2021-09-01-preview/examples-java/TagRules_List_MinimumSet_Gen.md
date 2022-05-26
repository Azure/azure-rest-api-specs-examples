Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dynatrace_1.0.0-beta.1/sdk/dynatrace/azure-resourcemanager-dynatrace/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for TagRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_List_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesListMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().list("myResourceGroup", "myMonitor", Context.NONE);
    }
}
```
