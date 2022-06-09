```java
import com.azure.core.util.Context;

/** Samples for Jobs Export. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/Common/TriggerExportJobs.json
     */
    /**
     * Sample code: Export Jobs.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void exportJobs(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.jobs().exportWithResponse("NetSDKTestRsVault", "SwaggerTestRg", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
