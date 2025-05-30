
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2ARemoveDisksInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RemoveDisksInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RemoveDisksInputProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationProtectedItems RemoveDisks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectedItems_RemoveDisks.json
     */
    /**
     * Sample code: Removes disk(s).
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        removesDiskS(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().removeDisks("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            new RemoveDisksInput()
                .withProperties(new RemoveDisksInputProperties().withProviderSpecificDetails(new A2ARemoveDisksInput()
                    .withVmDisksUris(Arrays.asList("https://vmstorage.blob.core.windows.net/vhds/datadisk1.vhd")))),
            com.azure.core.util.Context.NONE);
    }
}
