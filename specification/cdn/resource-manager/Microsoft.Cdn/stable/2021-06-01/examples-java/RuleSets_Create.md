```java
import com.azure.core.util.Context;

/** Samples for RuleSets Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/RuleSets_Create.json
     */
    /**
     * Sample code: RuleSets_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ruleSetsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRuleSets()
            .createWithResponse("RG", "profile1", "ruleSet1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
