
import com.azure.resourcemanager.azurestackhci.models.EdgeMachineRemoteSupportJobProperties;
import com.azure.resourcemanager.azurestackhci.models.RemoteSupportAccessLevel;
import com.azure.resourcemanager.azurestackhci.models.RemoteSupportType;
import java.time.OffsetDateTime;

/**
 * Samples for EdgeMachineJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_CreateOrUpdate_RemoteSupport.json
     */
    /**
     * Sample code: EdgeMachineJobs_CreateOrUpdate_RemoteSupport.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineJobsCreateOrUpdateRemoteSupport(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().define("RemoteSupport").withExistingEdgeMachine("ArcInstance-rg", "machine1")
            .withProperties(
                new EdgeMachineRemoteSupportJobProperties().withAccessLevel(RemoteSupportAccessLevel.DIAGNOSTICS)
                    .withExpirationTimestamp(OffsetDateTime.parse("2024-01-29T10:43:27.9471574Z"))
                    .withType(RemoteSupportType.ENABLE))
            .create();
    }
}
