import com.azure.resourcemanager.recoveryservicessiterecovery.models.TestFailoverCleanupInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.TestFailoverCleanupInputProperties;

/** Samples for ReplicationProtectedItems TestFailoverCleanup. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectedItems_TestFailoverCleanup.json
     */
    /**
     * Sample code: Execute test failover cleanup.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeTestFailoverCleanup(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationProtectedItems()
            .testFailoverCleanup(
                "vault1",
                "resourceGroupPS1",
                "cloud1",
                "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
                "f8491e4f-817a-40dd-a90c-af773978c75b",
                new TestFailoverCleanupInput()
                    .withProperties(new TestFailoverCleanupInputProperties().withComments("Test Failover Cleanup")),
                com.azure.core.util.Context.NONE);
    }
}
