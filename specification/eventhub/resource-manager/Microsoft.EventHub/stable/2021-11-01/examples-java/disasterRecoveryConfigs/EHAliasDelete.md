```java
import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasDelete.json
     */
    /**
     * Sample code: EHAliasDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eHAliasDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getDisasterRecoveryConfigs()
            .deleteWithResponse(
                "exampleResourceGroup", "sdk-Namespace-5849", "sdk-DisasterRecovery-3814", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
