```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.DataMoveLevel;
import com.azure.resourcemanager.recoveryservicesbackup.models.PrepareDataMoveRequest;

/** Samples for ResourceProvider BmsPrepareDataMove. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/BackupDataMove/PrepareDataMove_Post.json
     */
    /**
     * Sample code: Prepare Data Move.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void prepareDataMove(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceProviders()
            .bmsPrepareDataMove(
                "source-rsv",
                "sourceRG",
                new PrepareDataMoveRequest()
                    .withTargetResourceId(
                        "/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/targetRG/providers/Microsoft.RecoveryServices/vaults/target-rsv")
                    .withTargetRegion("USGov Virginia")
                    .withDataMoveLevel(DataMoveLevel.VAULT),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
