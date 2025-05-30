
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AAddDisksInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AVmDiskInputDetails;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.AddDisksInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.AddDisksInputProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationProtectedItems AddDisks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectedItems_AddDisks.json
     */
    /**
     * Sample code: Add disk(s) for protection.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        addDiskSForProtection(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().addDisks("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            new AddDisksInput().withProperties(new AddDisksInputProperties().withProviderSpecificDetails(
                new A2AAddDisksInput().withVmDisks(Arrays.asList(new A2AVmDiskInputDetails()
                    .withDiskUri("https://vmstorage.blob.core.windows.net/vhds/datadisk1.vhd")
                    .withRecoveryAzureStorageAccountId(
                        "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/recoveryResource/providers/Microsoft.Storage/storageAccounts/recoverystorage")
                    .withPrimaryStagingAzureStorageAccountId(
                        "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/primaryResource/providers/Microsoft.Storage/storageAccounts/vmcachestorage"))))),
            com.azure.core.util.Context.NONE);
    }
}
