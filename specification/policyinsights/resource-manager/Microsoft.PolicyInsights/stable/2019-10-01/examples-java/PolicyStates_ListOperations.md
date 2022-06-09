```java
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_ListOperations.json
     */
    /**
     * Sample code: List operations.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listOperations(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.
