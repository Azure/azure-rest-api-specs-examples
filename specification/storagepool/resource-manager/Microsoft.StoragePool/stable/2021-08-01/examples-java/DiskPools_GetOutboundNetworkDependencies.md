Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagepool_1.0.0-beta.1/sdk/storagepool/azure-resourcemanager-storagepool/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DiskPools ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_GetOutboundNetworkDependencies.json
     */
    /**
     * Sample code: Get Disk Pool outbound network dependencies.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void getDiskPoolOutboundNetworkDependencies(
        com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager
            .diskPools()
            .listOutboundNetworkDependenciesEndpoints("Sample-WestUSResourceGroup", "SampleAse", Context.NONE);
    }
}
```
