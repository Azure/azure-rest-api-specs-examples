Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-baremetalinfrastructure_1.0.0-beta.1/sdk/baremetalinfrastructure/azure-resourcemanager-baremetalinfrastructure/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalOperations_List.json
     */
    /**
     * Sample code: List all management operations supported by the AzureBareMetal RP.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void listAllManagementOperationsSupportedByTheAzureBareMetalRP(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager.operations().list(Context.NONE);
    }
}
```
