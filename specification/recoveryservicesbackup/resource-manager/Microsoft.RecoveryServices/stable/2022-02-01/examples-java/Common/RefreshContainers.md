```java
import com.azure.core.util.Context;

/** Samples for ProtectionContainers Refresh. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/Common/RefreshContainers.json
     */
    /**
     * Sample code: Trigger Azure Vm Discovery.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void triggerAzureVmDiscovery(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainers()
            .refreshWithResponse("NetSDKTestRsVault", "SwaggerTestRg", "Azure", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.5/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
