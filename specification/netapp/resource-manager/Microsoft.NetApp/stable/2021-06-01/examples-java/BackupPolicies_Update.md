Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.6/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.BackupPolicy;

/** Samples for BackupPolicies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/BackupPolicies_Update.json
     */
    /**
     * Sample code: BackupPolicies_Update.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupPoliciesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        BackupPolicy resource =
            manager.backupPolicies().getWithResponse("myRG", "account1", "backupPolicyName", Context.NONE).getValue();
        resource
            .update()
            .withDailyBackupsToKeep(5)
            .withWeeklyBackupsToKeep(10)
            .withMonthlyBackupsToKeep(10)
            .withEnabled(false)
            .apply();
    }
}
```
