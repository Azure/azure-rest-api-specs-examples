
import com.azure.resourcemanager.datatransfer.models.Direction;
import com.azure.resourcemanager.datatransfer.models.ListApprovedSchemasRequest;

/**
 * Samples for AzureDataTransfer ListApprovedSchemas.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/postListApprovedSchemas.json
     */
    /**
     * Sample code: Performs action request.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void performsActionRequest(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.azureDataTransfers().listApprovedSchemasWithResponse(
            new ListApprovedSchemasRequest().withPipeline("test").withDirection(Direction.SEND),
            com.azure.core.util.Context.NONE);
    }
}
