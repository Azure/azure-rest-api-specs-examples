Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.5/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.MoveRPAcrossTiersRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointTierType;

/** Samples for ResourceProvider MoveRecoveryPoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/TriggerRecoveryPointMove_Post.json
     */
    /**
     * Sample code: Trigger RP Move Operation.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void triggerRPMoveOperation(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceProviders()
            .moveRecoveryPoint(
                "testVault",
                "netsdktestrg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "348916168024334",
                new MoveRPAcrossTiersRequest()
                    .withObjectType("MoveRPAcrossTiersRequest")
                    .withSourceTierType(RecoveryPointTierType.HARDENED_RP)
                    .withTargetTierType(RecoveryPointTierType.ARCHIVED_RP),
                Context.NONE);
    }
}
```
