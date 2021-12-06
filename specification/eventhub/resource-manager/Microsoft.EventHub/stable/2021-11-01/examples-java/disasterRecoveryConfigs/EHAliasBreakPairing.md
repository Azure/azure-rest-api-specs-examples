Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs BreakPairing. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasBreakPairing.json
     */
    /**
     * Sample code: EHAliasBreakPairing.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eHAliasBreakPairing(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getDisasterRecoveryConfigs()
            .breakPairingWithResponse(
                "exampleResourceGroup", "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", Context.NONE);
    }
}
```
