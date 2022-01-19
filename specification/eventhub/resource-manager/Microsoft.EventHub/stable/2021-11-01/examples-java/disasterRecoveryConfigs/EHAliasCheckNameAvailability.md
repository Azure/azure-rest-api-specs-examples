Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.models.CheckNameAvailabilityParameter;

/** Samples for DisasterRecoveryConfigs CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasCheckNameAvailability.json
     */
    /**
     * Sample code: NamespacesCheckNameAvailability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void namespacesCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getDisasterRecoveryConfigs()
            .checkNameAvailabilityWithResponse(
                "exampleResourceGroup",
                "sdk-Namespace-9080",
                new CheckNameAvailabilityParameter().withName("sdk-DisasterRecovery-9474"),
                Context.NONE);
    }
}
```
