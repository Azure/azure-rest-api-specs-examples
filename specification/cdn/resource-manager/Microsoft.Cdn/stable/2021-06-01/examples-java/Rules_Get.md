```java
import com.azure.core.util.Context;

/** Samples for Rules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Get.json
     */
    /**
     * Sample code: Rules_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rulesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRules()
            .getWithResponse("RG", "profile1", "ruleSet1", "rule1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
