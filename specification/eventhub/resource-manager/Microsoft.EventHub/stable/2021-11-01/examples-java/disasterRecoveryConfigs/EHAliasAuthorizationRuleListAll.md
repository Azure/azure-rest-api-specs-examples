Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs ListAuthorizationRules. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleListAll.json
     */
    /**
     * Sample code: ListAuthorizationRules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAuthorizationRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getDisasterRecoveryConfigs()
            .listAuthorizationRules(
                "exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4047", Context.NONE);
    }
}
```
