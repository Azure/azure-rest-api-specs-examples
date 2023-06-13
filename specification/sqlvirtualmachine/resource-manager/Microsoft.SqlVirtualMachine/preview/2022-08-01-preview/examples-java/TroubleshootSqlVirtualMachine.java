import com.azure.resourcemanager.sqlvirtualmachine.fluent.models.SqlVmTroubleshootingInner;
import com.azure.resourcemanager.sqlvirtualmachine.models.TroubleshootingAdditionalProperties;
import com.azure.resourcemanager.sqlvirtualmachine.models.TroubleshootingScenario;
import com.azure.resourcemanager.sqlvirtualmachine.models.UnhealthyReplicaInfo;
import java.time.OffsetDateTime;

/** Samples for SqlVirtualMachineTroubleshoot Troubleshoot. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/TroubleshootSqlVirtualMachine.json
     */
    /**
     * Sample code: Start SQL virtual machine troubleshooting operation.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void startSQLVirtualMachineTroubleshootingOperation(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager
            .sqlVirtualMachineTroubleshoots()
            .troubleshoot(
                "testrg",
                "testvm",
                new SqlVmTroubleshootingInner()
                    .withStartTimeUtc(OffsetDateTime.parse("2022-07-09T17:10:00Z"))
                    .withEndTimeUtc(OffsetDateTime.parse("2022-07-09T22:10:00Z"))
                    .withTroubleshootingScenario(TroubleshootingScenario.UNHEALTHY_REPLICA)
                    .withProperties(
                        new TroubleshootingAdditionalProperties()
                            .withUnhealthyReplicaInfo(new UnhealthyReplicaInfo().withAvailabilityGroupName("AG1"))),
                com.azure.core.util.Context.NONE);
    }
}
