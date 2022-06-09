```java
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureIaaSvmProtectionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.DailyRetentionSchedule;
import com.azure.resourcemanager.recoveryservicesbackup.models.LongTermRetentionPolicy;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDuration;
import com.azure.resourcemanager.recoveryservicesbackup.models.RetentionDurationType;
import com.azure.resourcemanager.recoveryservicesbackup.models.ScheduleRunType;
import com.azure.resourcemanager.recoveryservicesbackup.models.SimpleSchedulePolicy;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for ProtectionPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Simple.json
     */
    /**
     * Sample code: Create or Update Simple Azure Vm Protection Policy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createOrUpdateSimpleAzureVmProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionPolicies()
            .define("testPolicy1")
            .withRegion((String) null)
            .withExistingVault("NetSDKTestRsVault", "SwaggerTestRg")
            .withProperties(
                new AzureIaaSvmProtectionPolicy()
                    .withSchedulePolicy(
                        new SimpleSchedulePolicy()
                            .withScheduleRunFrequency(ScheduleRunType.DAILY)
                            .withScheduleRunTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T02:00:00Z"))))
                    .withRetentionPolicy(
                        new LongTermRetentionPolicy()
                            .withDailySchedule(
                                new DailyRetentionSchedule()
                                    .withRetentionTimes(Arrays.asList(OffsetDateTime.parse("2018-01-24T02:00:00Z")))
                                    .withRetentionDuration(
                                        new RetentionDuration()
                                            .withCount(1)
                                            .withDurationType(RetentionDurationType.DAYS))))
                    .withTimeZone("Pacific Standard Time"))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.5/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
