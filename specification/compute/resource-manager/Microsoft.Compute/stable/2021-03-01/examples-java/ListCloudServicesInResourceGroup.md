Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for CloudServices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServicesInResourceGroup.json
     */
    /**
     * Sample code: List Cloud Services in a Resource Group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServicesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServices()
            .listByResourceGroup("ConstosoRG", Context.NONE);
    }
}
```
