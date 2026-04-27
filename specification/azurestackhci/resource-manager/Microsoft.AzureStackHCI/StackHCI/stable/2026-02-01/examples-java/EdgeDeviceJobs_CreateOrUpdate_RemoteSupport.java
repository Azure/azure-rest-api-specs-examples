
import com.azure.resourcemanager.azurestackhci.models.HciEdgeDeviceJob;
import com.azure.resourcemanager.azurestackhci.models.HciRemoteSupportJobProperties;
import com.azure.resourcemanager.azurestackhci.models.RemoteSupportAccessLevel;
import com.azure.resourcemanager.azurestackhci.models.RemoteSupportType;
import java.time.OffsetDateTime;

/**
 * Samples for EdgeDeviceJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/EdgeDeviceJobs_CreateOrUpdate_RemoteSupport.json
     */
    /**
     * Sample code: EdgeDeviceJobs_CreateOrUpdate_RemoteSupport.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeDeviceJobsCreateOrUpdateRemoteSupport(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeDeviceJobs().createOrUpdate(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1",
            "default", "collectLog",
            new HciEdgeDeviceJob().withProperties(
                new HciRemoteSupportJobProperties().withAccessLevel(RemoteSupportAccessLevel.DIAGNOSTICS)
                    .withExpirationTimestamp(OffsetDateTime.parse("2024-01-29T10:43:27.9471574Z"))
                    .withType(RemoteSupportType.ENABLE)),
            com.azure.core.util.Context.NONE);
    }
}
