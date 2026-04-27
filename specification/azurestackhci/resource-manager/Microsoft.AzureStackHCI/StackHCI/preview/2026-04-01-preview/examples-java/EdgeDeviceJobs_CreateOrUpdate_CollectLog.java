
import com.azure.resourcemanager.azurestackhci.models.HciCollectLogJobProperties;
import com.azure.resourcemanager.azurestackhci.models.HciEdgeDeviceJob;
import java.time.OffsetDateTime;

/**
 * Samples for EdgeDeviceJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeDeviceJobs_CreateOrUpdate_CollectLog.json
     */
    /**
     * Sample code: EdgeDeviceJobs_CreateOrUpdate_CollectLog.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeDeviceJobsCreateOrUpdateCollectLog(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeDeviceJobs().createOrUpdate(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1",
            "default", "collectLog",
            new HciEdgeDeviceJob().withProperties(
                new HciCollectLogJobProperties().withFromDate(OffsetDateTime.parse("2024-01-29T10:43:27.9471574Z"))
                    .withToDate(OffsetDateTime.parse("2024-01-29T10:43:27.9471574Z"))),
            com.azure.core.util.Context.NONE);
    }
}
