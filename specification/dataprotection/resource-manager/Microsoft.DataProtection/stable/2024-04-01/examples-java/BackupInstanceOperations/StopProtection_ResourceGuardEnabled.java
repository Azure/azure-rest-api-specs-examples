
import com.azure.resourcemanager.dataprotection.models.StopProtectionRequest;
import java.util.Arrays;

/**
 * Samples for BackupInstances StopProtection.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/StopProtection_ResourceGuardEnabled.json
     */
    /**
     * Sample code: StopProtection with MUA.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void stopProtectionWithMUA(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().stopProtection("testrg", "testvault", "testbi",
            new StopProtectionRequest().withResourceGuardOperationRequests(Arrays.asList(
                "/subscriptions/754ec39f-8d2a-44cf-bfbf-13107ac85c36/resourcegroups/mua-testing/providers/Microsoft.DataProtection/resourceGuards/gvjreddy-test-ecy-rg-reader/dppDisableStopProtectionRequests/default")),
            com.azure.core.util.Context.NONE);
    }
}
