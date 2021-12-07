Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyMetadata GetResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyMetadata_GetResource.json
     */
    /**
     * Sample code: Get a single policy metadata resource.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getASinglePolicyMetadataResource(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyMetadatas().getResourceWithResponse("NIST_SP_800-53_R4_AC-2", Context.NONE);
    }
}
```
