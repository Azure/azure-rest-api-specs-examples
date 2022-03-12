Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyExemptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForSubscription.json
     */
    /**
     * Sample code: List policy exemptions that apply to a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicyExemptionsThatApplyToASubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyExemptions().list("atScope()", Context.NONE);
    }
}
```
