
import com.azure.resourcemanager.datatransfer.models.ListPipelineMetadataRequest;

/**
 * Samples for AzureDataTransfer ListPipelineMetadata.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/AzureDataTransfer_ListPipelineMetadata.json
     */
    /**
     * Sample code: AzureDataTransfer_ListPipelineMetadata.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        azureDataTransferListPipelineMetadata(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.azureDataTransfers().listPipelineMetadataWithResponse(
            new ListPipelineMetadataRequest().withPipeline("testPipeline"), com.azure.core.util.Context.NONE);
    }
}
