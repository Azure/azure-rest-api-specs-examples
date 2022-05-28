Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Policies GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyGet.json
     */
    /**
     * Sample code: Get Policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getPolicies()
            .getByResourceGroupWithResponse("rg1", "MicrosoftCdnWafPolicy", Context.NONE);
    }
}
```
