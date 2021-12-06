Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ResourceGuardProxyOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/ResourceGuardProxyCRUD/GetResourceGuardProxy.json
     */
    /**
     * Sample code: Get ResourceGuardProxy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getResourceGuardProxy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceGuardProxyOperations()
            .getWithResponse("sampleVault", "SampleResourceGroup", "swaggerExample", Context.NONE);
    }
}
```
