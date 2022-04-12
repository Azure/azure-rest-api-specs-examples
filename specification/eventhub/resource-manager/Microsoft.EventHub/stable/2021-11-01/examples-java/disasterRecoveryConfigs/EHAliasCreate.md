Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.fluent.models.ArmDisasterRecoveryInner;

/** Samples for DisasterRecoveryConfigs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasCreate.json
     */
    /**
     * Sample code: EHAliasCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eHAliasCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getDisasterRecoveryConfigs()
            .createOrUpdateWithResponse(
                "exampleResourceGroup",
                "sdk-Namespace-8859",
                "sdk-DisasterRecovery-3814",
                new ArmDisasterRecoveryInner().withPartnerNamespace("sdk-Namespace-37"),
                Context.NONE);
    }
}
```
