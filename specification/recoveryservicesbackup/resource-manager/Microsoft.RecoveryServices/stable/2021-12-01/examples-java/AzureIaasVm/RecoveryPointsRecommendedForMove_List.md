Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.3/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.ListRecoveryPointsRecommendedForMoveRequest;
import java.util.Arrays;

/** Samples for RecoveryPointsRecommendedForMove List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/RecoveryPointsRecommendedForMove_List.json
     */
    /**
     * Sample code: Get Protected Azure Vm Recovery Points Recommended for Move.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedAzureVmRecoveryPointsRecommendedForMove(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .recoveryPointsRecommendedForMoves()
            .list(
                "rshvault",
                "rshhtestmdvmrg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
                "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
                new ListRecoveryPointsRecommendedForMoveRequest()
                    .withObjectType("ListRecoveryPointsRecommendedForMoveRequest")
                    .withExcludedRPList(Arrays.asList("348916168024334", "348916168024335")),
                Context.NONE);
    }
}
```
