Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hardwaresecuritymodules_1.0.0-beta.1/sdk/hardwaresecuritymodules/azure-resourcemanager-hardwaresecuritymodules/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DedicatedHsm ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/GetOutboundNetworkDependenciesEndpointsList.json
     */
    /**
     * Sample code: List OutboundNetworkDependenciesEndpoints by Managed Cluster.
     *
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void listOutboundNetworkDependenciesEndpointsByManagedCluster(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().listOutboundNetworkDependenciesEndpoints("hsm-group", "hsm1", Context.NONE);
    }
}
```
