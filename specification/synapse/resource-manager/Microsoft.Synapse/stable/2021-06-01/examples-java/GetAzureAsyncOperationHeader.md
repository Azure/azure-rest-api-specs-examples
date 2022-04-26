Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Operations GetAzureAsyncHeaderResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetAzureAsyncOperationHeader.json
     */
    /**
     * Sample code: Get azure async operation header result.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAzureAsyncOperationHeaderResult(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .operations()
            .getAzureAsyncHeaderResultWithResponse(
                "resourceGroup1", "workspace1", "01234567-89ab-4def-0123-456789abcdef", Context.NONE);
    }
}
```
